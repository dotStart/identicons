package color

import (
	"github.com/lucasb-eyer/go-colorful"
	"math/bits"
)

type hslGenerator struct {
	saturation float64
	lightness  float64
}

func HSL(saturation float64, lightness float64) Generator {
	return &hslGenerator{saturation, lightness}
}

func (*hslGenerator) RequiredBits() uint {
	return uint(bits.Len(uint(359)))
}

func (g *hslGenerator) Generate(code uint64) colorful.Color {
	hue := float64(code % 360)
	return colorful.Hsl(hue, g.saturation, g.lightness)
}
