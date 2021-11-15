package tiled

import (
	svg "github.com/ajstarks/svgo"
	"github.com/dotstart/identicons/library/identicons/shape"
)

type shapeTile struct {
	shape *shape.Shape2d
}

// Shape creates a new tile which draws a given shape consisting of two or more vertices.
func Shape(vertices ...*shape.Vert2d) Tile {
	return &shapeTile{shape.New(vertices...)}
}

func (t *shapeTile) Draw(canvas *svg.SVG, x float64, y float64, size float64, flipX bool, flipY bool, rotations uint) {
	s := *t.shape

	rotations = rotations % 4 // 4 rotations have no effect
	if rotations > 0 {
		for i := uint(0); i < rotations; i++ {
			s.Rotate()
		}
	}

	s.Flip(flipX, flipY)
	s.Draw(canvas, x, y, size, size, nil, nil)
}
