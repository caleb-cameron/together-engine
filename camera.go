package engine

import (
	"math"

	"github.com/faiface/pixel"
)

type Camera struct {
	Position  pixel.Vec
	Speed     float64
	Zoom      float64
	ZoomSpeed float64
	Matrix    pixel.Matrix
}

func NewCamera() *Camera {
	c := &Camera{}

	return c
}

func (c *Camera) UpdateZoom(zoom float64) {
	c.Zoom *= math.Pow(c.ZoomSpeed, zoom)
}
