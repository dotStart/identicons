package tiled

import (
	svg "github.com/ajstarks/svgo"
	"math"
)

type circleTile struct {
	scale float64
}

// Circle returns a fully circular tile with a given scale.
func Circle(scale float64) Tile {
	return &circleTile{scale}
}

func (t *circleTile) Draw(canvas *svg.SVG, x float64, y float64, size float64, flipX bool, flipY bool, rotations uint) {
	r := size * 0.5

	ra := int(math.Floor(r * t.scale))
	cx := int(math.Floor(x + r))
	cy := int(math.Floor(y + r))

	canvas.Circle(cx, cy, ra)
}
