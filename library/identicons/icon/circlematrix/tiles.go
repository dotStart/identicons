package circlematrix

import (
	"github.com/dotstart/identicons/library/identicons/icon/tiled"
)

var tileTable = []tiled.Tile{
	// filled circles
	tiled.Circle(1),
	tiled.Circle(0.75),
	tiled.Circle(0.5),
	tiled.Circle(0.25),

	// bold donuts
	tiled.Donut(1, 0.5),
	tiled.Donut(0.75, 0.25),

	// lightweight donuts
	tiled.Donut(1, 0.75),
	tiled.Donut(0.75, 0.5),
	tiled.Donut(0.5, 0.25),
}
