package color

import (
	"github.com/lucasb-eyer/go-colorful"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHsvGenerator_Generate(t *testing.T) {
	gen := HSV(1, 1)

	assert.Equal(t, gen.Generate(0), colorful.Hsv(0, 1, 1))
	assert.Equal(t, gen.Generate(90), colorful.Hsv(90, 1, 1))
	assert.Equal(t, gen.Generate(180), colorful.Hsv(180, 1, 1))
	assert.Equal(t, gen.Generate(270), colorful.Hsv(270, 1, 1))
	assert.Equal(t, gen.Generate(360), colorful.Hsv(0, 1, 1))
	assert.Equal(t, gen.Generate(450), colorful.Hsv(90, 1, 1))
}
