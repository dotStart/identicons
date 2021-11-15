package block

import (
	"github.com/dotstart/identicons/internal/util"
	"github.com/dotstart/identicons/library/identicons/icon/tiled"
)

var halfRectTile = tiled.Rect(0, 0, 0.5, 0.5)

var tileTable = []tiled.Tile{
	tiled.Rect(0, 0, 1, 1), // filled tile
	halfRectTile,           // quarter tile
	tiled.Combined(halfRectTile, tiled.Flipped(halfRectTile, true, true)), // diagonally connected rects

	tiled.Shape(util.Vert(0, 0), util.Vert(0, 1), util.Vert(0.25, 1), util.Vert(0.25, 0.25), util.Vert(1, 0.25), util.Vert(1, 0)),                                                                                       // corner
	tiled.Shape(util.Vert(0, 0), util.Vert(0, 0.25), util.Vert(0.375, 0.25), util.Vert(0.375, 1), util.Vert(0.625, 1), util.Vert(0.625, 0.25), util.Vert(1, 0.25), util.Vert(1, 0)),                                     // T
	tiled.Shape(util.Vert(0, 0), util.Vert(0, 1), util.Vert(1, 1), util.Vert(1, 0), util.Vert(0.25, 0), util.Vert(0.25, 0.25), util.Vert(0.75, 0.25), util.Vert(0.75, 0.75), util.Vert(0.25, 0.75), util.Vert(0.25, 0)), // Hollow Rect
}
