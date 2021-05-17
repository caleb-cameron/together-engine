package engine

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"

	quadtree "github.com/JamesLMilner/quadtree-go"

	perlin "github.com/aquilax/go-perlin"
	"github.com/faiface/pixel"
)

type Chunk struct {
	Bounds    pixel.Rect
	Tiles     [][]Tile // Rows then columns
	Generated bool
	changed   bool
	mutex     sync.RWMutex
	Quadtree  *quadtree.Quadtree
}

var Noise *perlin.Perlin

func NewChunk(bounds pixel.Rect) *Chunk {
	columns := make([][]Tile, int(bounds.W()))
	for col := range columns {
		columns[col] = make([]Tile, int(bounds.H()))
	}

	return &Chunk{
		Bounds: bounds,
		Tiles:  columns,
		mutex:  sync.RWMutex{},
		Quadtree: &quadtree.Quadtree{
			Bounds: quadtree.Bounds{
				X:      bounds.Min.X,
				Y:      bounds.Min.Y,
				Width:  bounds.W(),
				Height: bounds.H(),
			},
			MaxObjects: EntityQuadtreeMaxLevels,
			MaxLevels:  EntityQuadtreeMaxObjectsPerNode,
			Level:      0,
			Objects:    make([]quadtree.Bounds, 0),
			Nodes:      make([]quadtree.Quadtree, 0),
		},
	}
}

func (c *Chunk) Lock() {
	c.mutex.Lock()
}

func (c *Chunk) Unlock() {
	c.mutex.Unlock()
}

func (c *Chunk) RLock() {
	c.mutex.RLock()
}

func (c *Chunk) RUnlock() {
	c.mutex.RUnlock()
}

func ChunkCanBeLoaded(x, y int) bool {
	fileInfo, err := os.Stat(fmt.Sprintf("%s%d_%d", ChunksDir, x, y))

	if err != nil {
		if os.IsNotExist(err) {
			return false
		}

		panic(err)
	}

	return !fileInfo.IsDir()
}

func LoadChunk(x, y int) *Chunk {
	//fmt.Printf("Loading chunk (%d, %d)\n", x, y)
	var chunkData []byte

	f, err := os.Open(fmt.Sprintf("%s%d_%d", ChunksDir, x, y))

	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}

		panic(err)
	}
	defer f.Close()

	chunkData, err = ioutil.ReadAll(f)

	if err != nil {
		panic(err)
	}

	if len(chunkData) == 0 {
		return nil
	}

	c, err := DecodeChunk(chunkData)

	if err != nil {
		fmt.Printf("Error reading Chunk: %d,%d\n", x, y)
		panic(err)
	}

	return c
}

func (c *Chunk) GetChunkPos() pixel.Vec {
	if c == nil {
		log.Println("Attempt to call GetChunkPos on nil chunk")
		return pixel.Vec{}
	}
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return pixel.Vec{
		X: c.Bounds.Min.X / float64(ChunkSize),
		Y: c.Bounds.Min.Y / float64(ChunkSize),
	}
}

func (c *Chunk) GetNeighborChunkPositions() []pixel.Vec {
	if c == nil {
		return nil
	}

	c.mutex.RLock()
	defer c.mutex.RUnlock()

	pos := c.GetChunkPos()
	return []pixel.Vec{
		{pos.X, pos.Y + 1},
		{pos.X + 1, pos.Y + 1},
		{pos.X + 1, pos.Y},
		{pos.X + 1, pos.Y - 1},
		{pos.X, pos.Y - 1},
		{pos.X - 1, pos.Y - 1},
		{pos.X - 1, pos.Y},
		{pos.X - 1, pos.Y + 1},
	}
}

func (c *Chunk) Generate() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for x, col := range c.Tiles {
		xpos := int(c.Bounds.Min.X) + x
		for y, _ := range col {
			ypos := int(c.Bounds.Min.Y) + y
			altitudeNoise := Noise.Noise2D(float64(xpos)*0.01, float64(ypos)*0.01)
			tileHeight := int(float64(WorldMaxAltitude)*altitudeNoise) + 10
			//fmt.Printf("tileHeight: %d\n", tileHeight)
			//fmt.Println(altitudeNoise)
			if tileHeight < 5 {
				c.Tiles[x][y] = Tiles[TileWater]
			} else {
				c.Tiles[x][y] = Tiles[TileGrass]
			}

			c.Tiles[x][y].Chunk = c
		}
	}

	c.Generated = true
}

func (c *Chunk) GenerateBoundaryTiles() {
	chunkPos := c.GetChunkPos()

	c.mutex.Lock()
	defer c.mutex.Unlock()

	xpos, ypos := int(chunkPos.X), int(chunkPos.Y)
	var (
		leftChunk        *Chunk
		rightChunk       *Chunk
		topChunk         *Chunk
		bottomChunk      *Chunk
		topRightChunk    *Chunk
		topLeftChunk     *Chunk
		bottomRightChunk *Chunk
		bottomLeftChunk  *Chunk

		// thisTile              *Tile
		leftTile              *Tile
		rightTile             *Tile
		topTile               *Tile
		bottomTile            *Tile
		topRightCornerTile    *Tile
		topLeftCornerTile     *Tile
		bottomRightCornerTile *Tile
		bottomLeftCornerTile  *Tile
	)

	hasLeftChunk := GWorld.ChunkExists(xpos-1, ypos)
	hasRightChunk := GWorld.ChunkExists(xpos+1, ypos)
	hasTopChunk := GWorld.ChunkExists(xpos, ypos+1)
	hasBottomChunk := GWorld.ChunkExists(xpos, ypos-1)
	hasTopRightChunk := GWorld.ChunkExists(xpos+1, ypos+1)
	hasTopLeftChunk := GWorld.ChunkExists(xpos-1, ypos+1)
	hasBottomRightChunk := GWorld.ChunkExists(xpos+1, ypos-1)
	hasBottomLeftChunk := GWorld.ChunkExists(xpos-1, ypos-1)

	if hasLeftChunk {
		leftChunk = GWorld.GetChunk(xpos-1, ypos)
	}
	if hasRightChunk {
		rightChunk = GWorld.GetChunk(xpos+1, ypos)
	}
	if hasTopChunk {
		topChunk = GWorld.GetChunk(xpos, ypos+1)
	}
	if hasBottomChunk {
		bottomChunk = GWorld.GetChunk(xpos, ypos-1)
	}
	if hasBottomLeftChunk {
		bottomLeftChunk = GWorld.GetChunk(xpos-1, ypos-1)
	}
	if hasBottomRightChunk {
		bottomRightChunk = GWorld.GetChunk(xpos+1, ypos-1)
	}
	if hasTopLeftChunk {
		topLeftChunk = GWorld.GetChunk(xpos-1, ypos+1)
	}
	if hasTopRightChunk {
		topRightChunk = GWorld.GetChunk(xpos+1, ypos+1)
	}

	for x, col := range c.Tiles {
		for y, _ := range col {
			// thisTile = &c.Tiles[x][y]
			if x == 0 {
				// This Tile is along the left edge
				if hasLeftChunk {
					leftTile = &leftChunk.Tiles[ChunkSize-1][y]
				}
				if y == 0 {
					// This Tile is at the bottom left corner
					topRightCornerTile = &c.Tiles[x+1][y+1]

					if hasBottomLeftChunk {
						bottomLeftCornerTile = &bottomLeftChunk.Tiles[ChunkSize-1][ChunkSize-1]
					}
					if hasBottomChunk {
						bottomRightCornerTile = &bottomChunk.Tiles[x+1][ChunkSize-1]
					}
					if hasLeftChunk {
						topLeftCornerTile = &leftChunk.Tiles[ChunkSize-1][y+1]
					}
				} else if y == ChunkSize-1 {
					// This Tile is at the top left corner
					bottomRightCornerTile = &c.Tiles[x+1][y-1]

					if hasTopLeftChunk {
						topLeftCornerTile = &topLeftChunk.Tiles[ChunkSize-1][0]
					}
					if hasLeftChunk {
						bottomLeftCornerTile = &leftChunk.Tiles[ChunkSize-1][y-1]
					}
					if hasTopChunk {
						topRightCornerTile = &topChunk.Tiles[x+1][0]
					}
				} else if y < ChunkSize-1 && y > 0 {
					// This Tile is along the left edge but not a corner
					if hasLeftChunk {
						topLeftCornerTile = &leftChunk.Tiles[ChunkSize-1][y+1]
						bottomLeftCornerTile = &leftChunk.Tiles[ChunkSize-1][y-1]
					}
					bottomRightCornerTile = &c.Tiles[x+1][y-1]
					topRightCornerTile = &c.Tiles[x+1][y+1]
				}

			} else {
				// This Tile is not on the left edge
				// so we can just use its neighbor Tile
				leftTile = &c.Tiles[x-1][y]
			}
			if x == ChunkSize-1 {
				// This Tile is along the right edge
				if hasRightChunk && rightChunk.Generated {
					rightTile = &rightChunk.Tiles[0][y]
				}

				if y == 0 {
					// This Tile is on the bottom right corner
					topLeftCornerTile = &c.Tiles[x-1][y+1]
					if hasBottomRightChunk {
						bottomRightCornerTile = &bottomRightChunk.Tiles[0][0]
					}
					if hasBottomChunk {
						bottomLeftCornerTile = &bottomChunk.Tiles[x-1][ChunkSize-1]
					}
					if hasRightChunk {
						topRightCornerTile = &rightChunk.Tiles[0][0]
					}
				} else if y == ChunkSize-1 {
					// This Tile is on the top right corner
					if hasTopRightChunk && topRightChunk.Tiles != nil {
						topRightCornerTile = &topRightChunk.Tiles[0][0]
					}
					if hasRightChunk {
						bottomRightCornerTile = &rightChunk.Tiles[0][y-1]
					}
					if hasTopChunk {
						topLeftCornerTile = &topChunk.Tiles[y-1][0]
					}
					bottomLeftCornerTile = &c.Tiles[x-1][y-1]
				} else if y < ChunkSize-1 && y > 0 {
					// This Tile is along the right edge but not a corner
					if hasRightChunk {
						topRightCornerTile = &rightChunk.Tiles[0][y+1]
						bottomRightCornerTile = &rightChunk.Tiles[0][y-1]
					}
					bottomLeftCornerTile = &c.Tiles[x-1][y-1]
					topLeftCornerTile = &c.Tiles[x-1][y+1]
				}
			} else {
				// This Tile is not on the right edge,
				// so we can just use its neighbor Tile.
				rightTile = &c.Tiles[x+1][y]
			}

			if x > 0 && x < ChunkSize-1 {
				// This Tile is not on either end of the X axis
				if y == 0 {
					// This Tile is along the bottom edge,  but not a corner
					if hasBottomChunk {
						bottomTile = &bottomChunk.Tiles[x][ChunkSize-1]
						bottomRightCornerTile = &bottomChunk.Tiles[x+1][ChunkSize-1]
						bottomLeftCornerTile = &bottomChunk.Tiles[x-1][ChunkSize-1]
					}

					topRightCornerTile = &c.Tiles[x+1][y+1]
					topLeftCornerTile = &c.Tiles[x-1][y+1]
				} else if y == ChunkSize-1 {
					// This Tile is along the top edge, but not a corner
					bottomTile = &c.Tiles[x][y-1]
					if hasTopChunk {
						topTile = &topChunk.Tiles[x][0]
						topRightCornerTile = &topChunk.Tiles[x+1][0]
						topLeftCornerTile = &topChunk.Tiles[x-1][0]
					}

					bottomRightCornerTile = &c.Tiles[x+1][y-1]
					bottomLeftCornerTile = &c.Tiles[x-1][y-1]
				} else {
					// This Tile is not on any edge.
					topRightCornerTile = &c.Tiles[x+1][y+1]
					topLeftCornerTile = &c.Tiles[x-1][y+1]
					bottomRightCornerTile = &c.Tiles[x+1][y-1]
					bottomLeftCornerTile = &c.Tiles[x-1][y-1]
				}
			}

			if y == ChunkSize-1 {
				if hasTopChunk {
					topTile = &topChunk.Tiles[x][0]
				}
			} else {
				topTile = &c.Tiles[x][y+1]
			}

			if y == 0 {
				if hasBottomChunk {
					bottomTile = &bottomChunk.Tiles[x][ChunkSize-1]
				}
			} else {
				bottomTile = &c.Tiles[x][y-1]
			}

			c.Tiles[x][y].LeftNeighbor = leftTile
			c.Tiles[x][y].RightNeighbor = rightTile
			c.Tiles[x][y].TopNeighbor = topTile
			c.Tiles[x][y].BottomNeighbor = bottomTile
			c.Tiles[x][y].TopLeftNeighbor = topLeftCornerTile
			c.Tiles[x][y].TopRightNeighbor = topRightCornerTile
			c.Tiles[x][y].BottomLeftNeighbor = bottomRightCornerTile
			c.Tiles[x][y].BottomRightNeighbor = bottomLeftCornerTile

			c.Tiles[x][y].DecideTileType()
		}
	}

}

func (c *Chunk) GetTile(x, y int) *Tile {
	c.RLock()
	defer c.RUnlock()

	return &c.Tiles[x][y]
}

func (c *Chunk) ReplaceTile(x, y int, tile Tile) {
	c.Lock()
	defer c.Unlock()

	c.Tiles[x][y] = tile
	c.changed = true
}

func (c *Chunk) PersistToDisk() {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	b, err := c.Encode()

	if err != nil {
		panic(err)
	}

	ioutil.WriteFile(fmt.Sprintf("%s%d_%d", ChunksDir, int(c.GetChunkPos().X), int(c.GetChunkPos().Y)), b, 0666)
	c.changed = false
}

func (c *Chunk) Encode() ([]byte, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return GWorld.Encoder.Encode(c)
}

func DecodeChunk(b []byte) (*Chunk, error) {
	return GWorld.Encoder.Decode(b)
}
