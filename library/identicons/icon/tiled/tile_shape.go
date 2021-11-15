package tiled

import (
	svg "github.com/ajstarks/svgo"
	"github.com/dotstart/identicons/internal/util"
)

type shapeTile struct {
	shape *util.Shape
}

// Shape creates a new tile which draws a given shape consisting of two or more vertices.
func Shape(vertices ...*util.Vert2d) Tile {
	return &shapeTile{util.NewShape(vertices...)}
}

func (t *shapeTile) Draw(canvas *svg.SVG, x float64, y float64, size float64, flipX bool, flipY bool, rotations uint) {
	shape := *t.shape

	rotations = rotations % 4 // 4 rotations have no effect
	if rotations > 0 {
		for i := uint(0); i < rotations; i++ {
			shape.Rotate()
		}
	}

	shape.Flip(flipX, flipY)
	shape.Draw(canvas, x, y, size, size, nil, nil)
}
