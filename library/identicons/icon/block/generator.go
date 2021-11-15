package block

import (
	"github.com/dotstart/identicons/library/identicons/icon"
	"github.com/dotstart/identicons/library/identicons/icon/tiled"
)

// New creates a new generator which relies on a matrix of circles.
func New(opts ...tiled.Option) icon.Generator {
	return tiled.New("block", tileTable, tileTable, tileTable, opts...)
}
