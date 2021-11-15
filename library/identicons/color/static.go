package color

import "github.com/lucasb-eyer/go-colorful"

type staticGenerator struct {
	color colorful.Color
}

// Static returns a color generator which always returns the same value regardless of the input
// code.
func Static(color colorful.Color) Generator {
	return &staticGenerator{color}
}

func (*staticGenerator) RequiredBits() uint {
	return 0
}

func (g *staticGenerator) Generate(code uint64) colorful.Color {
	return g.color
}
