package modern

import (
	"github.com/dotstart/identicons/library/identicons/icon/tiled"
	"github.com/dotstart/identicons/library/identicons/shape"
)

var halfRectangleTile = tiled.Rect(0, 0, 0.5, 0.5)

var tileTable = []tiled.Tile{
	halfRectangleTile, // half rectangle
	tiled.Combined(halfRectangleTile, tiled.Flipped(halfRectangleTile, true, true)), // diagonally connected half rectangles

	tiled.Shape(shape.Vert(0, 0), shape.Vert(1, 0), shape.Vert(1, 1)), // triangle down-left
	tiled.Shape(shape.Vert(0, 1), shape.Vert(1, 1), shape.Vert(1, 0)), // triangle center
	tiled.Shape(shape.Vert(0, 0), shape.Vert(1, 1), shape.Vert(1, 0)), // triangle top-right
	tiled.Shape(shape.Vert(0, 0), shape.Vert(0, 1), shape.Vert(1, 0)), // triangle out

	tiled.Shape(shape.Vert(0, 0), shape.Vert(0, 1), shape.Vert(0.5, 0)), // steep triangle down
	tiled.Shape(shape.Vert(0, 0.5), shape.Vert(0, 1), shape.Vert(1, 1)), // steep triangle right
	tiled.Shape(shape.Vert(0.5, 1), shape.Vert(1, 1), shape.Vert(1, 0)), // steep triangle up
	tiled.Shape(shape.Vert(0, 0), shape.Vert(1, 0.5), shape.Vert(1, 0)), // steep triangle left

	tiled.Shape(shape.Vert(.5, 0), shape.Vert(0, .5), shape.Vert(.5, 1), shape.Vert(1, .5)), // diamond

	tiled.Circle(1),
	tiled.Circle(0.5),
	tiled.Donut(1, 0.75),
	tiled.Donut(0.5, 0.25),
}

var centerTileTable = append(
	tileTable,

	tiled.Shape(shape.Vert(0, 0), shape.Vert(0, 1), shape.Vert(0.5, 1), shape.Vert(0.25, 0.25), shape.Vert(1, 0.5), shape.Vert(1, 0)),                                                                                            // inverse star
	tiled.Shape(shape.Vert(0, 0), shape.Vert(0, 1), shape.Vert(1, 1), shape.Vert(1, 0), shape.Vert(0.25, 0), shape.Vert(0.25, 0.25), shape.Vert(0.75, 0.25), shape.Vert(0.5, 0.75), shape.Vert(0.25, 0.25), shape.Vert(0.25, 0)), // inverse triangle
	tiled.Shape(shape.Vert(0, 0), shape.Vert(0, 1), shape.Vert(1, 1), shape.Vert(1, 0), shape.Vert(0.27, 0.2), shape.Vert(0.8, 0.73), shape.Vert(0.73, 0.8), shape.Vert(0.2, 0.27), shape.Vert(0.27, 0.2), shape.Vert(0.27, 0)),  // inverse diagonal bar
)
