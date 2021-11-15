package classic

import (
	"github.com/dotstart/identicons/library/identicons/icon"
	"github.com/dotstart/identicons/library/identicons/icon/tiled"
)

// New creates a new classic generator which is heavily inspired by the original identicon
// implementation by Don Park (https://github.com/donpark/identicon).
func New(opts ...tiled.Option) icon.Generator {
	return tiled.New("classic", tileTable, tileTable, centerTileTable, opts...)
}
