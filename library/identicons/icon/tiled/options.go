package tiled

import (
	"github.com/dotstart/identicons/library/identicons/code"
	"github.com/dotstart/identicons/library/identicons/color"
	"github.com/lucasb-eyer/go-colorful"
	"math"
)

// Option represents an option which may be passed to New in order to adjust the default settings
// of the generator.
type Option func(g *Generator)

// Code selects the code generator which shall be used to create a seed code from input values.
func Code(generator code.Generator) Option {
	return func(g *Generator) {
		g.codeGenerator = generator
	}
}

// ForegroundColor selects a generator which decides the color which is to be applied to the entire
// image or its respective sections.
func ForegroundColor(color color.Generator) Option {
	return func(g *Generator) {
		g.foregroundColorGenerator = color
	}
}

// BackgroundColor selects the color of the background rectangle of the resulting image.
func BackgroundColor(color *colorful.Color) Option {
	return func(g *Generator) {
		g.backgroundColor = color
	}
}

// TransparentBackground disables the background coloring for the resulting image thus making it
// transparent where tiles did not draw geometry.
func TransparentBackground() Option {
	return BackgroundColor(nil)
}

// ColoredSides causes the unique sections within the image to be colored differently to make
// them easier to distinguish.
func ColoredSides(value bool) Option {
	return func(g *Generator) {
		g.colorSides = value
	}
}

// TileSize customizes the default size at which each respective tile within the graphic is
// rendered.
func TileSize(tileSize uint) Option {
	return func(g *Generator) {
		g.tileSize = tileSize
	}
}

// ImageSize selects a suitable TileSize based on the given input image size.
//
// All image sizes must be divisible by four in order to produce accurate results.
func ImageSize(imageSize uint) Option {
	tileSize := math.Floor(float64(imageSize) / float64(tileCount))
	return TileSize(uint(tileSize))
}

// PermitAdjacentDuplicates selects whether the same tile may appear in adjacent areas within the
// resulting image.
func PermitAdjacentDuplicates(value bool) Option {
	return func(g *Generator) {
		g.permitAdjacentDuplicates = value
	}
}
