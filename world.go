package engine

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"sync"

	"github.com/faiface/pixel"
)

type World struct {
	Chunks      map[int]map[int]*Chunk
	Encoder     *Encoder
	updateMutex *sync.RWMutex
}

func (w *World) Init() {
	w.Encoder = NewEncoder()
	w.updateMutex = new(sync.RWMutex)

	err := os.Mkdir(DataDir, 0700)

	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	err = os.Mkdir(ChunksDir, 0700)

	if err != nil && !os.IsExist(err) {
		panic(err)
	}
}

func (w *World) Generate() {

	numChunksH := 2 //(int(gWindow.Bounds().H()) / TileSize / ChunkSize) + 1
	numChunksW := 2 //(int(gWindow.Bounds().W()) / TileSize / ChunkSize) + 1

	w.Chunks = make(map[int]map[int]*Chunk, numChunksW)

	for x := 0 - numChunksW/2; x < numChunksW/2; x++ {
		w.Chunks[x] = make(map[int]*Chunk, numChunksH)
		for y := 0 - numChunksH/2; y < numChunksH/2; y++ {
			if ChunkCanBeLoaded(x, y) {
				w.LoadChunk(x, y)
			} else {
				w.CreateChunk(x, y)
			}
		}
	}

	for _, row := range w.Chunks {
		if row == nil {
			row = make(map[int]*Chunk, 10)
		}
		for _, c := range row {
			if !c.Generated {
				//fmt.Printf("Generating c %d, %d\n", x, y)
				c.Generate()
			}
		}
	}
}

func (w *World) CreateChunkAtPos(pos pixel.Vec) {
	w.CreateChunk(int(pos.X)/ChunkSize/TileSize, int(pos.Y)/ChunkSize/TileSize)
}

func (w *World) CreateChunk(x, y int) {

	if w.ChunkExists(x, y) {
		return
	} else if ChunkCanBeLoaded(x, y) {
		w.LoadChunk(x, y)
		return
	}

	c := NewChunk(pixel.R(float64(x*ChunkSize), float64(y*ChunkSize), float64((x+1)*ChunkSize), float64(y+1)*ChunkSize))
	//fmt.Printf("Generated Chunk %+v at %d, %d.\n", c.GetChunkPos(), x, y)
	c.Generate()
	c.GenerateBoundaryTiles()
	c.PersistToDisk()

	w.updateMutex.Lock()
	if _, ok := w.Chunks[x]; !ok {
		w.Chunks[x] = make(map[int]*Chunk, 10)
	}
	w.Chunks[x][y] = c
	w.updateMutex.Unlock()
}

func (w *World) UpdateLoadedChunks() {
	playerX := GPlayer.Position.X / ChunkSize / TileSize
	playerY := GPlayer.Position.Y / ChunkSize / TileSize

	chunkLoadPadding := 2.0

	loadRect := pixel.R(playerX-ChunkLoadRadius-chunkLoadPadding, playerY-ChunkLoadRadius-chunkLoadPadding, playerX+ChunkLoadRadius+chunkLoadPadding, playerY+ChunkLoadRadius+chunkLoadPadding)

	for x, col := range w.Chunks {
		for y, _ := range col {
			if !loadRect.Contains(pixel.V(float64(x), float64(y))) {
				delete(w.Chunks[x], y)
			}
		}

		if len(w.Chunks[x]) == 0 {
			delete(w.Chunks, x)
		}
	}
}

func (w *World) LoadChunk(x, y int) {
	if _, ok := w.Chunks[x]; !ok {
		w.Chunks[x] = make(map[int]*Chunk, 10)
	}

	w.updateMutex.Lock()
	if w.Chunks == nil {
		w.Chunks = make(map[int]map[int]*Chunk, 10)
	}

	if _, ok := w.Chunks[x]; !ok {
		w.Chunks[x] = make(map[int]*Chunk, 10)
	}

	w.Chunks[x][y] = LoadChunk(x, y)
	w.updateMutex.Unlock()
}

func (w *World) UnloadChunk(x, y int) {
	c := w.GetChunk(x, y)

	if c == nil {
		return
	}

	if c.changed {
		c.PersistToDisk()
	}

	w.updateMutex.Lock()
	delete(w.Chunks[x], y)
	w.updateMutex.Unlock()
}

func (w *World) PreloadNeighborChunks(c *Chunk) {
	neighbors := c.GetNeighborChunkPositions()
	for _, pos := range neighbors {
		x := int(pos.X)
		y := int(pos.Y)

		w.LoadOrCreateChunk(x, y)

		c := w.GetChunk(x, y)
		if c != nil {
			// Load a bigger ring around this Chunk
			nextNeighbors := c.GetNeighborChunkPositions()
			for _, npos := range nextNeighbors {
				x = int(npos.X)
				y = int(npos.Y)

				w.LoadOrCreateChunk(x, y)
			}
		}
	}
}

func (w *World) LoadOrCreateChunk(x, y int) {
	if !w.ChunkExists(x, y) {
		if ChunkCanBeLoaded(x, y) {
			w.LoadChunk(x, y)
		} else {
			w.CreateChunk(x, y)
		}
	}
}

func (w *World) GenerateBoundaryTiles() {
	for _, row := range w.Chunks {
		for _, chunk := range row {
			chunk.GenerateBoundaryTiles()
		}
	}
}

func (w *World) Draw() {
	w.updateMutex.RLock()
	for _, col := range w.Chunks {
		for _, c := range col {
			c.Draw(pixel.IM.Moved(pixel.Vec{X: c.Bounds.Min.X * TileSize, Y: c.Bounds.Min.Y * TileSize}))
		}
	}
	w.updateMutex.RUnlock()
}

func (w *World) ChunkExists(x int, y int) bool {
	w.updateMutex.RLock()
	defer w.updateMutex.RUnlock()
	if _, ok := w.Chunks[x]; !ok {
		return false
	}
	_, ok := w.Chunks[x][y]
	return ok
}

func (w *World) GetChunk(x, y int) *Chunk {
	w.updateMutex.RLock()
	defer w.updateMutex.RUnlock()

	if _, ok := w.Chunks[x]; !ok {
		return nil
	}

	if _, ok := w.Chunks[x][y]; !ok {
		return nil
	}

	return w.Chunks[x][y]
}

func (w *World) GetChunkForPos(pos pixel.Vec) *Chunk {
	x := int(pos.X) / ChunkSize / TileSize
	y := int(pos.Y) / ChunkSize / TileSize

	if w.ChunkExists(x, y) {
		return w.GetChunk(x, y)
	}

	return nil
}

func (w *World) SetSeed(seed uint64) {
	seedFilePath := fmt.Sprintf("%s/seed", DataDir)
	seedStr := strconv.Itoa(int(seed))

	err := ioutil.WriteFile(seedFilePath, []byte(seedStr), 0600)
	if err != nil {
		panic(err)
	}
}

func (w *World) GetSeed() uint64 {
	seedFilePath := fmt.Sprintf("%s/seed", DataDir)
	if !fileExists(seedFilePath) {
		return 0
	}

	f, err := os.Open(seedFilePath)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	bytesData, err := ioutil.ReadAll(f)

	if err != nil {
		panic(err)
	}

	seed, err := strconv.Atoi(string(bytesData))

	return uint64(seed)

}
