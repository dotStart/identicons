package modern

import (
	"github.com/dotstart/identicons/library/identicons/icon"
	"github.com/dotstart/identicons/library/identicons/icon/tiled"
)

// New creates a new icon generator which is similar to the classic implementation but provides some
// more complex tiles as well as colorful sides by default.
func New(opts ...tiled.Option) icon.Generator {
	o := []tiled.Option{
		tiled.ColoredSides(true),
	}

	return tiled.New("modern", tileTable, tileTable, centerTileTable, append(o, opts...)...)
}
