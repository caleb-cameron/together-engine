package engine

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/aquilax/go-perlin"
)

func TestSerializeChunk(t *testing.T) {
	ChunksDir = "data/chunks/"
	ChunkSize = 10
	ChunkLoadRadius = 1
	WorldMaxAltitude = 30

	GWorld = World{}
	GWorld.Init()

	seed := uint64(time.Now().UnixNano())
	GWorld.SetSeed(seed)

	Noise = perlin.NewPerlin(2.0, 2.0, 3, int64(seed))

	GWorld.Generate()
	GWorld.GenerateBoundaryTiles()

	c1 := GWorld.Chunks[0][0]

	serialized := SerializeChunk(*c1)

	c2 := DeserializeChunk(serialized)

	same, err := CompareChunks(c1, c2)

	if !same {
		t.Errorf(err.Error())
	}

}

func CompareChunks(c1 *Chunk, c2 *Chunk) (bool, error) {
	if !c1.Bounds.Min.Eq(c2.Bounds.Min) {
		return false, errors.New(fmt.Sprintf("min bounds not equal: before: %+v, after: %+v", c1.Bounds.Min, c2.Bounds.Min))
	}

	if !c1.Bounds.Max.Eq(c2.Bounds.Max) {
		return false, errors.New(fmt.Sprintf("max bounds not equal: before: %+v, after: %+v", c1.Bounds.Max, c2.Bounds.Max))
	}

	if c1.Bounds.W() != c2.Bounds.W() {
		return false, errors.New(fmt.Sprintf("chunk width not equal: before: %+v, after: %+v", c1.Bounds.W(), c2.Bounds.W()))
	}

	if c1.Bounds.H() != c2.Bounds.H() {
		return false, errors.New(fmt.Sprintf("chunk height not equal: before: %+v, after: %+v", c1.Bounds.H(), c2.Bounds.H()))
	}

	for y := range c1.Tiles {
		for x := range c1.Tiles {
			if c1.Tiles[y][x].Id != c2.Tiles[y][x].Id {
				return false, errors.New(fmt.Sprintf("tiles at (%d, %d) not equal: before: %+v, after: %+v", y, x, c1.Tiles[y][x], c2.Tiles[y][x]))
			}
		}
	}

	return true, nil
}
