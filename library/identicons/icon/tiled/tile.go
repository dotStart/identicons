package tiled

import (
	svg "github.com/ajstarks/svgo"
)

// Tile represents an arbitrary tile implementation which renders at the given desired location
type Tile interface {
	Draw(canvas *svg.SVG, x float64, y float64, size float64, flipX bool, flipY bool, rotations uint)
}
