package tiled

import (
	"fmt"
	svg "github.com/ajstarks/svgo"
	"math"
)

type donutTile struct {
	outerScale float64
	innerScale float64
}

// Donut created a tile which draws a circle with a hole in the middle.
func Donut(outerScale float64, innerScale float64) Tile {
	return &donutTile{outerScale, innerScale}
}

func (t *donutTile) Draw(canvas *svg.SVG, x float64, y float64, size float64, flipX bool, flipY bool, rotations uint) {
	r := size * 0.5
	s := int(math.Floor(size))

	sx := int(math.Floor(x))
	sy := int(math.Floor(y))

	cx := int(math.Floor(x + r))
	cy := int(math.Floor(y + r))

	ro := int(math.Floor(r * t.outerScale))
	ri := int(math.Floor(r * t.innerScale))

	maskId := fmt.Sprintf("donut_mask_%dX%d", cx, cy)
	canvas.Mask(maskId, 0, 0, s, s, "maskContentUnits=\"userSpaceOnUse\"")
	{
		canvas.Rect(sx, sy, s, s, "fill=\"white\"")
		canvas.Circle(cx, cy, ri, "fill=\"black\"")
	}
	canvas.MaskEnd()

	canvas.Circle(cx, cy, ro, fmt.Sprintf("mask=\"url(#%s)\"", maskId))
}
