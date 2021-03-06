package engine

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"sync"

	"github.com/faiface/pixel"
)

type World struct {
	Chunks      map[int]map[int]*Chunk
	Encoder     *Encoder
	UpdateMutex *sync.RWMutex
}

func (w *World) Init() {
	w.Encoder = NewEncoder()
	w.UpdateMutex = new(sync.RWMutex)

	err := os.Mkdir(DataDir, 0700)

	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	err = os.Mkdir(ChunksDir, 0700)

	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	w.Chunks = make(map[int]map[int]*Chunk, 10)

	for i, row := range w.Chunks {
		if row == nil {
			w.Chunks[i] = make(map[int]*Chunk, 10)
		}
	}
}

func (w *World) Generate() {
	numChunksH := 2 //(int(gWindow.Bounds().H()) / TileSize / ChunkSize) + 1
	numChunksW := 2 //(int(gWindow.Bounds().W()) / TileSize / ChunkSize) + 1

	for x := 0 - numChunksW/2; x < numChunksW/2; x++ {
		for y := 0 - numChunksH/2; y < numChunksH/2; y++ {
			if ChunkCanBeLoaded(x, y) {
				w.LoadChunk(x, y)
			} else {
				w.CreateChunk(x, y)
			}
		}
	}

	for _, row := range w.Chunks {
		for _, c := range row {
			if !c.Generated {
				// fmt.Printf("Generating c %d, %d\n", int(c.Bounds.Min.X)/ChunkSize, int(c.Bounds.Min.Y)/ChunkSize)
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

	c := NewChunk(pixel.R(float64(x*ChunkSize), float64(y*ChunkSize), float64((x+1)*ChunkSize), float64(y+1)*float64(ChunkSize)))
	//fmt.Printf("Generated Chunk %+v at %d, %d.\n", c.GetChunkPos(), x, y)
	c.Generate()
	c.GenerateBoundaryTiles()
	c.PersistToDisk()

	// Regenerate boundary tiles for neighbor chunks now that this one is created.
	neighborChunks := c.GetNeighborChunkPositions()

	for _, pos := range neighborChunks {
		n := GWorld.GetChunk(int(pos.X), int(pos.Y))

		if n != nil {
			n.GenerateBoundaryTiles()
			n.PersistToDisk()
		}
	}

	w.UpdateMutex.Lock()
	if _, ok := w.Chunks[x]; !ok {
		w.Chunks[x] = make(map[int]*Chunk, 10)
	}
	w.Chunks[x][y] = c
	w.UpdateMutex.Unlock()
}
func (w *World) UpdateLoadedChunksServer() {
	players := PlayerList.GetPlayers()

	w.UpdateLoadedChunks(players)
}

func (w *World) GetPlayerLoadRect(player *Player) pixel.Rect {
	playerX := player.GetPosition().X / float64(ChunkSize) / float64(TileSize)
	playerY := player.GetPosition().Y / float64(ChunkSize) / float64(TileSize)

	loadRect := pixel.R(playerX-ChunkLoadRadius-ChunkLoadPadding, playerY-ChunkLoadRadius-ChunkLoadPadding, playerX+ChunkLoadRadius+ChunkLoadPadding, playerY+ChunkLoadRadius+ChunkLoadPadding)

	return loadRect
}

func (w *World) GetChunksToLoadForPlayer(player *Player) []pixel.Vec {
	loadRect := w.GetPlayerLoadRect(player)

	return getChunksCoordsInRect(loadRect)
}

func (w *World) UpdateLoadedChunks(players map[string]*Player) {
	keepList := []pixel.Vec{}
	w.UpdateMutex.RLock()

	for _, player := range players {
		keepList = append(keepList, w.GetChunksToLoadForPlayer(player)...)
	}

	w.UpdateMutex.RUnlock()

	w.PruneChunks(keepList)

	for _, pos := range keepList {
		x := int(pos.X)
		y := int(pos.Y)

		if !w.ChunkExists(x, y) {
			w.LoadOrCreateChunk(x, y)
		}
	}
}

func (w *World) InsertChunk(x int, y int, c *Chunk) {
	// w.UpdateMutex.Lock()
	// defer w.UpdateMutex.Unlock()

	if _, ok := w.Chunks[x]; !ok {
		w.Chunks[x] = map[int]*Chunk{}
	}

	w.Chunks[x][y] = c
}

func (w *World) PruneChunks(keepList []pixel.Vec) {
	// w.UpdateMutex.Lock()
	// defer w.UpdateMutex.Unlock()

	for x, col := range w.Chunks {
		for y := range col {
			found := false

			for _, c := range keepList {
				if int(c.X) == x && int(c.Y) == y {
					found = true
					break
				}
			}

			if !found {
				w.Chunks[x][y].Lock()
				delete(w.Chunks[x], y)
			}
		}
		if len(w.Chunks[x]) == 0 {
			w.UpdateMutex.Lock()
			delete(w.Chunks, x)
			w.UpdateMutex.Unlock()
		}
	}
}

func (w *World) CanLoadChunk(x, y int) bool {
	return ChunkCanBeLoaded(x, y)
}

func (w *World) GetNumLoadedChunks() int {
	w.UpdateMutex.RLock()
	defer w.UpdateMutex.RUnlock()

	out := 0

	for _, row := range w.Chunks {
		out += len(row)
	}

	return out
}

func (w *World) GetLoadedChunks() map[int]map[int]*Chunk {
	w.UpdateMutex.RLock()
	defer w.UpdateMutex.RUnlock()

	return w.Chunks
}

func (w *World) LoadChunk(x, y int) {
	if _, ok := w.Chunks[x]; !ok {
		w.Chunks[x] = make(map[int]*Chunk, 10)
	}

	w.UpdateMutex.Lock()
	if w.Chunks == nil {
		w.Chunks = make(map[int]map[int]*Chunk, 10)
	}

	if _, ok := w.Chunks[x]; !ok {
		w.Chunks[x] = make(map[int]*Chunk, 10)
	}

	w.Chunks[x][y] = LoadChunk(x, y)
	w.UpdateMutex.Unlock()
}

func (w *World) UnloadChunk(x, y int) {
	c := w.GetChunk(x, y)

	if c == nil {
		return
	}

	if c.changed {
		c.PersistToDisk()
	}

	w.UpdateMutex.Lock()
	delete(w.Chunks[x], y)
	w.UpdateMutex.Unlock()
}

func (w *World) PreloadNeighborChunks(c *Chunk) {
	if c == nil {
		log.Println("Tried to preload neighbors of nil chunk.")
		return
	}
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

func (w *World) ChunkExists(x int, y int) bool {
	w.UpdateMutex.RLock()
	defer w.UpdateMutex.RUnlock()
	if _, ok := w.Chunks[x]; !ok {
		return false
	}
	_, ok := w.Chunks[x][y]
	return ok
}

func (w *World) GetChunk(x, y int) *Chunk {
	w.UpdateMutex.RLock()
	defer w.UpdateMutex.RUnlock()

	if _, ok := w.Chunks[x]; !ok {
		return nil
	}

	if _, ok := w.Chunks[x][y]; !ok {
		return nil
	}

	return w.Chunks[x][y]
}

func (w *World) GamePosToChunkAndTilePos(pos pixel.Vec) (int, int, int, int) {
	tileXOffset := pos.X / float64(TileSize)
	tileYOffset := pos.Y / float64(TileSize)

	chunkXOffset := tileXOffset / float64(ChunkSize)
	chunkYOffset := tileYOffset / float64(ChunkSize)

	chunkX := int(math.Floor(chunkXOffset))
	chunkY := int(math.Floor(chunkYOffset))

	tileX := int(math.Floor(tileXOffset)) % ChunkSize
	tileY := int(math.Floor(tileYOffset)) % ChunkSize

	if tileX < 0 {
		tileX += ChunkSize
	}

	if tileY < 0 {
		tileY += ChunkSize
	}

	return chunkX, chunkY, tileX, tileY
}

func (w *World) GetChunkForPos(pos pixel.Vec) *Chunk {
	chunkX, chunkY, _, _ := w.GamePosToChunkAndTilePos(pos)

	if w.ChunkExists(chunkX, chunkY) {
		return w.GetChunk(chunkX, chunkY)
	}

	return nil
}

func (w *World) GetTileForPos(pos pixel.Vec) *Tile {
	tileX := int(pos.X) / ChunkSize % TileSize
	tileY := int(pos.Y) / ChunkSize % TileSize

	chunkX := (int(pos.X) - (ChunkSize / 2)) / ChunkSize / TileSize
	chunkY := (int(pos.Y) - (ChunkSize / 2)) / ChunkSize / TileSize

	if tileX < 0 {
		chunkX -= 1
		tileX += ChunkSize
	}

	if tileY < 0 {
		chunkY -= 1
		tileY += ChunkSize
	}

	if !w.ChunkExists(chunkX, chunkY) {
		return nil
	}

	c := w.GetChunk(chunkX, chunkY)

	if c == nil {
		return nil
	}

	return c.GetTile(tileX, tileY)
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
