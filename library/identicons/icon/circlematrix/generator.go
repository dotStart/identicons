package circlematrix

import (
	"github.com/dotstart/identicons/library/identicons/icon"
	"github.com/dotstart/identicons/library/identicons/icon/tiled"
)

// New creates a new generator which relies on a matrix of circles.
func New(opts ...tiled.Option) icon.Generator {
	o := []tiled.Option{
		tiled.ColoredSides(true),
	}

	return tiled.New("circle-matrix", tileTable, tileTable, tileTable, append(o, opts...)...)
}
