package color

import (
	"github.com/lucasb-eyer/go-colorful"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStaticGenerator_Generate(t *testing.T) {
	expected := colorful.LinearRgb(0.5, 0.5, 0.5)
	gen := Static(expected)

	for i := 0; i < 64; i++ {
		actual := gen.Generate(uint64(i * 7))

		assert.Equal(t, expected, actual)
	}
}
