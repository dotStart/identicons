package tiled

import (
	svg "github.com/ajstarks/svgo"
)

type transformedTile struct {
	tile Tile

	offsetX   float64
	offsetY   float64
	scale     float64
	rotations uint
}

// Transform adjusts the position and scale at which a given tile is drawn.
func Transform(tile Tile, offsetX float64, offsetY float64, scale float64, rotations uint) Tile {
	return &transformedTile{tile, offsetX, offsetY, scale, rotations}
}

func (t *transformedTile) Draw(canvas *svg.SVG, x float64, y float64, size float64, flipX bool, flipY bool, rotations uint) {
	ax := x + t.offsetX
	ay := y + t.offsetY
	as := size * t.scale
	r := (rotations + t.rotations) % 4

	t.tile.Draw(canvas, ax, ay, as, flipX, flipY, r)
}
