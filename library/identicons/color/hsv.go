package color

import "github.com/lucasb-eyer/go-colorful"

type hsvGenerator struct {
	saturation float64
	value      float64
}

func HSV(saturation float64, value float64) Generator {
	return &hsvGenerator{saturation, value}
}

func (*hsvGenerator) RequiredBits() uint {
	return 9
}

func (g *hsvGenerator) Generate(code uint64) colorful.Color {
	hue := float64(code % 360)
	return colorful.Hsv(hue, g.saturation, g.value)
}
