package tiled

import (
	svg "github.com/ajstarks/svgo"
)

type flippedTile struct {
	tile Tile

	flipX bool
	flipY bool
}

// Flipped returns a tile which has been flipped along its X, Y axis or both.
func Flipped(tile Tile, flipX bool, flipY bool) Tile {
	return &flippedTile{tile, flipX, flipY}
}

func (t *flippedTile) Draw(canvas *svg.SVG, x float64, y float64, size float64, flipX bool, flipY bool, rotations uint) {
	fx := flipX != t.flipX
	fy := flipY != t.flipY

	t.tile.Draw(canvas, x, y, size, fx, fy, rotations)
}
