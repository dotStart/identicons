package classic

import (
	"github.com/dotstart/identicons/internal/util"
	"github.com/dotstart/identicons/library/identicons/icon/tiled"
)

var tileTable = []tiled.Tile{
	tiled.Shape(util.Vert(0, 0), util.Vert(1, 0), util.Vert(1, 1)), // triangle down-left
	tiled.Shape(util.Vert(0, 1), util.Vert(1, 1), util.Vert(1, 0)), // triangle center
	tiled.Shape(util.Vert(0, 0), util.Vert(1, 1), util.Vert(1, 0)), // triangle top-right
	tiled.Shape(util.Vert(0, 0), util.Vert(0, 1), util.Vert(1, 0)), // triangle out

	tiled.Shape(util.Vert(0, 0), util.Vert(0, 1), util.Vert(0.5, 0)), // steep triangle down
	tiled.Shape(util.Vert(0, 0.5), util.Vert(0, 1), util.Vert(1, 1)), // steep triangle right
	tiled.Shape(util.Vert(0.5, 1), util.Vert(1, 1), util.Vert(1, 0)), // steep triangle up
	tiled.Shape(util.Vert(0, 0), util.Vert(1, 0.5), util.Vert(1, 0)), // steep triangle left
}

var centerTileTable = append(
	tileTable,

	tiled.Shape(util.Vert(0, 0), util.Vert(0, 1), util.Vert(0.5, 1), util.Vert(0.25, 0.25), util.Vert(1, 0.5), util.Vert(1, 0)), // inverse star
)
