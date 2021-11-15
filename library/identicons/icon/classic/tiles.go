package classic

import (
	"github.com/dotstart/identicons/library/identicons/icon/tiled"
	"github.com/dotstart/identicons/library/identicons/shape"
)

var tileTable = []tiled.Tile{
	tiled.Shape(shape.Vert(0, 0), shape.Vert(1, 0), shape.Vert(1, 1)), // triangle down-left
	tiled.Shape(shape.Vert(0, 1), shape.Vert(1, 1), shape.Vert(1, 0)), // triangle center
	tiled.Shape(shape.Vert(0, 0), shape.Vert(1, 1), shape.Vert(1, 0)), // triangle top-right
	tiled.Shape(shape.Vert(0, 0), shape.Vert(0, 1), shape.Vert(1, 0)), // triangle out

	tiled.Shape(shape.Vert(0, 0), shape.Vert(0, 1), shape.Vert(0.5, 0)), // steep triangle down
	tiled.Shape(shape.Vert(0, 0.5), shape.Vert(0, 1), shape.Vert(1, 1)), // steep triangle right
	tiled.Shape(shape.Vert(0.5, 1), shape.Vert(1, 1), shape.Vert(1, 0)), // steep triangle up
	tiled.Shape(shape.Vert(0, 0), shape.Vert(1, 0.5), shape.Vert(1, 0)), // steep triangle left
}

var centerTileTable = append(
	tileTable,

	tiled.Shape(shape.Vert(0, 0), shape.Vert(0, 1), shape.Vert(0.5, 1), shape.Vert(0.25, 0.25), shape.Vert(1, 0.5), shape.Vert(1, 0)), // inverse star
)
