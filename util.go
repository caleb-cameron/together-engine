package engine

import (
	"os"

	"github.com/faiface/pixel"
)

func getChunksCoordsInRect(rect pixel.Rect) []pixel.Vec {

	out := []pixel.Vec{}

	for x := rect.Min.X - ChunkLoadPadding; x < rect.Max.X+1+ChunkLoadPadding; x++ {
		for y := rect.Min.Y - ChunkLoadPadding; y < rect.Max.Y+1+ChunkLoadPadding; y++ {
			out = append(out, pixel.Vec{X: x, Y: y})
		}
	}

	return out
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	var list []int
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func fileExists(path string) bool {
	fileInfo, err := os.Stat(path)

	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		panic(err)
	}

	return !fileInfo.IsDir()
}

func stringInSlice(s string, slice []string) bool {
	for _, s2 := range slice {
		if s == s2 {
			return true
		}
	}

	return false
}
