package engine

import (
	"encoding/binary"

	"github.com/faiface/pixel"
)

/*
	Serialization format:

	Chunk x (1 uint16)
	Chunk y (1 uint16)
	Chunk width (1 uint16)
	Chunk height (1 uint16)

	tile ids (1 uint16 * width * height)
*/

func SerializeChunk(c Chunk) []byte {
	w := uint16(c.Bounds.W())
	h := uint16(c.Bounds.H())
	x := uint16(c.Bounds.Min.X / ChunkSize)
	y := uint16(c.Bounds.Min.Y / ChunkSize)

	out := make([]byte, 4+(2*w*h)) // 4 bytes for x and y then 2 bytes for each tile
	curVal := make([]byte, 2)      // Placeholder for a little-endian representation of the current value.
	byteIndex := 0                 // keep track of our current place in the output byte slice

	binary.LittleEndian.PutUint16(curVal, x)
	out[byteIndex] = curVal[0]
	out[byteIndex+1] = curVal[1]
	byteIndex += 2

	binary.LittleEndian.PutUint16(curVal, y)
	out[byteIndex] = curVal[0]
	out[byteIndex+1] = curVal[1]
	byteIndex += 2

	for _, col := range c.Tiles {
		for _, tile := range col {
			binary.LittleEndian.PutUint16(curVal, uint16(tile.Id))
			out[byteIndex] = curVal[0]
			out[byteIndex+1] = curVal[1]
			byteIndex += 2
		}
	}

	return out
}

func DeserializeChunk(b []byte) *Chunk {
	x := int16(binary.LittleEndian.Uint16(b[0:2]))
	y := int16(binary.LittleEndian.Uint16(b[2:4]))

	c := NewChunk(pixel.R(float64(x*ChunkSize), float64(y*ChunkSize), float64((x+1)*ChunkSize), float64(y+1)*ChunkSize))

	byteIndex := 4
	for x := range c.Tiles {
		for y := range c.Tiles[x] {
			c.Tiles[x][y] = tiles[int(binary.LittleEndian.Uint16(b[byteIndex:byteIndex+2]))]
			byteIndex += 2
		}
	}

	return c
}
