package engine

import (
	"testing"
	"time"

	perlin "github.com/aquilax/go-perlin"
)

func init() {
	DataDir = "data/"
	ChunksDir = "data/chunks/"
	ChunkSize = 10
	ChunkLoadRadius = 1
	WorldMaxAltitude = 30

	GWorld = World{}
	GWorld.Init()
}

func TestGenerateWorld(t *testing.T) {
	seed := uint64(time.Now().UnixNano())
	GWorld.SetSeed(seed)

	Noise = perlin.NewPerlin(2.0, 2.0, 3, int64(seed))

	GWorld.Generate()
	GWorld.GenerateBoundaryTiles()

	for x := -1; x < 1; x++ {
		for y := -1; y < 1; y++ {
			if !GWorld.ChunkExists(x, y) {
				t.Errorf("Expected chunk %d,%d to exist after world generation.", x, y)
			}
		}
	}

	if t.Failed() {
		t.Logf("Chunks loaded: %+v\n", GWorld.GetLoadedChunks())
	}
}
