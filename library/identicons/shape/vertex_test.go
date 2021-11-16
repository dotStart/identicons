package shape

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVert(t *testing.T) {
	for x := 0; x < 16; x++ {
		for y := 0; y < 16; y++ {
			vert := Vert(float64(x), float64(y))

			assert.Equal(t, x, int(vert.X))
			assert.Equal(t, y, int(vert.Y))
		}
	}
}

func TestVert2d_Invert(t *testing.T) {
	for x := 0; x < 16; x++ {
		for y := 0; y < 16; y++ {
			vert := Vert(float64(x), float64(y))
			vert.Invert()

			assert.Equal(t, -x, int(vert.X))
			assert.Equal(t, -y, int(vert.Y))
		}
	}
}

func TestVert2d_Rotate(t *testing.T) {
	for x := 0; x < 16; x++ {
		for y := 0; y < 16; y++ {
			vert := Vert(float64(x), float64(y))
			vert.Rotate()

			assert.Equal(t, -y, int(vert.X))
			assert.Equal(t, x, int(vert.Y))
		}
	}
}

func TestVert2d_Multiply(t *testing.T) {
	for x := 0; x < 16; x++ {
		for y := 0; y < 16; y++ {
			vert := Vert(float64(x), float64(y))
			vert.Multiply(Vert(3, 7))

			assert.Equal(t, x*3, int(vert.X))
			assert.Equal(t, y*7, int(vert.Y))
		}
	}
}

func TestMinVert(t *testing.T) {
	a := Vert(-1, -1)
	b := Vert(0, 0)
	c := Vert(1, 1)
	d := Vert(-1, 1)
	e := Vert(1, -1)

	assert.Equal(t, Vert(-1, -1), MinVert(a, a))
	assert.Equal(t, Vert(-1, -1), MinVert(a, b))
	assert.Equal(t, Vert(-1, -1), MinVert(a, c))
	assert.Equal(t, Vert(-1, -1), MinVert(a, d))
	assert.Equal(t, Vert(-1, -1), MinVert(a, e))

	assert.Equal(t, Vert(0, 0), MinVert(b, c))
	assert.Equal(t, Vert(-1, 0), MinVert(b, d))
	assert.Equal(t, Vert(0, -1), MinVert(b, e))
}
