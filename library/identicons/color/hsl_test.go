package color

import (
	"github.com/lucasb-eyer/go-colorful"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHslGenerator_Generate(t *testing.T) {
	gen := HSL(1, 1)

	assert.Equal(t, gen.Generate(0), colorful.Hsl(0, 1, 1))
	assert.Equal(t, gen.Generate(90), colorful.Hsl(90, 1, 1))
	assert.Equal(t, gen.Generate(180), colorful.Hsl(180, 1, 1))
	assert.Equal(t, gen.Generate(270), colorful.Hsl(270, 1, 1))
	assert.Equal(t, gen.Generate(360), colorful.Hsl(0, 1, 1))
	assert.Equal(t, gen.Generate(450), colorful.Hsl(90, 1, 1))
}
