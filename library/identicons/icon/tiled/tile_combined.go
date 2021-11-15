package tiled

import (
	svg "github.com/ajstarks/svgo"
)

type combinedTile struct {
	tiles []Tile
}

// Combined returns a tile which renders multiple tiles within a single cell.
func Combined(tiles ...Tile) Tile {
	return &combinedTile{tiles}
}

func (t *combinedTile) Draw(canvas *svg.SVG, x float64, y float64, size float64, flipX bool, flipY bool, rotations uint) {
	for _, tile := range t.tiles {
		tile.Draw(canvas, x, y, size, flipX, flipY, rotations)
	}
}
