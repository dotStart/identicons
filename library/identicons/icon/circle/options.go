package circle

import (
	"github.com/dotstart/identicons/library/identicons/code"
	"github.com/dotstart/identicons/library/identicons/color"
	"github.com/lucasb-eyer/go-colorful"
)

// Option represents an option which may be passed to New in order to adjust the default settings
// of the generator.
type Option func(g *generator)

// Code selects the code generator which shall be used to create a seed code from input values.
func Code(code code.Generator) Option {
	return func(g *generator) {
		g.codeGenerator = code
	}
}

// ForegroundColor selects a generator which decides the color which is to be applied to the entire
// image or its respective sections.
func ForegroundColor(color color.Generator) Option {
	return func(g *generator) {
		g.foregroundColorGenerator = color
	}
}

// BackgroundColor selects the color of the background rectangle of the resulting image.
func BackgroundColor(color *colorful.Color) Option {
	return func(g *generator) {
		g.backgroundColor = color
	}
}

// TransparentBackground disables the background coloring for the resulting image thus making it
// transparent where tiles did not draw geometry.
func TransparentBackground() Option {
	return BackgroundColor(nil)
}

// Rings customizes the number of rings to draw within the image.
func Rings(rings uint) Option {
	return func(g *generator) {
		g.rings = rings
	}
}

// Segments customizes the number of segments to draw within the image.
func Segments(segments uint) Option {
	return func(g *generator) {
		g.segments = segments
	}
}

// DrawCore identifies whether the image core shall be drawn in addition to the rings.
func DrawCore(draw bool) Option {
	return func(g *generator) {
		g.drawCore = draw
	}
}

// RingWidth customizes the width of each ring.
func RingWidth(height uint) Option {
	return func(g *generator) {
		g.ringWidth = height
	}
}
