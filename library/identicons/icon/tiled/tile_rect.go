package tiled

import (
	svg "github.com/ajstarks/svgo"
	"github.com/dotstart/identicons/internal/util"
	"math"
)

var originPoints = []*util.Vert2d{
	util.Vert(-1, -1), // top-left
	util.Vert(1, -1),  // top-right
	util.Vert(1, 1),   // bottom-right
	util.Vert(-1, 1),  // bottom-left
}

type rectTile struct {
	offsetX float64
	offsetY float64

	scaleX float64
	scaleY float64
}

// Rect returns a tile which draws a simple rectangle.
func Rect(offsetX float64, offsetY float64, scaleX float64, scaleY float64) Tile {
	return &rectTile{offsetX, offsetY, scaleX, scaleY}
}

func (t *rectTile) Draw(canvas *svg.SVG, x float64, y float64, size float64, flipX bool, flipY bool, rotations uint) {
	width := size * t.scaleX
	height := size * t.scaleY
	origin := *originPoints[rotations%4]

	if flipX {
		origin.X *= -1
	}
	if flipY {
		origin.Y *= -1
	}

	destination := origin
	destination.Invert()
	destination.Multiply(util.Vert(t.scaleX, t.scaleY))

	origin.X = math.Max(0, origin.X)
	origin.Y = math.Max(0, origin.Y)
	destination.Plus(&origin)

	begin := util.MinVert(&origin, &destination)

	ax := int(math.Floor(x + (begin.X * size)))
	ay := int(math.Floor(y + (begin.Y * size)))
	w := int(math.Floor(width))
	h := int(math.Floor(height))

	canvas.Rect(ax, ay, w, h)
}
