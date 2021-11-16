package shape

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShape2d_Flip(t *testing.T) {
	s := Shape(
		Vert(0, 0),
		Vert(0.25, 0.75),
		Vert(0.75, 0.25),
		Vert(1, 1),
	)

	s.Flip(true, false)

	assert.Equal(t, Vert(1, 0), s.vertices[0])
	assert.Equal(t, Vert(0.75, 0.75), s.vertices[1])
	assert.Equal(t, Vert(0.25, 0.25), s.vertices[2])
	assert.Equal(t, Vert(0, 1), s.vertices[3])

	s.Flip(false, true)

	assert.Equal(t, Vert(1, 1), s.vertices[0])
	assert.Equal(t, Vert(0.75, 0.25), s.vertices[1])
	assert.Equal(t, Vert(0.25, 0.75), s.vertices[2])
	assert.Equal(t, Vert(0, 0), s.vertices[3])

	s.Flip(true, true)

	assert.Equal(t, Vert(0, 0), s.vertices[0])
	assert.Equal(t, Vert(0.25, 0.75), s.vertices[1])
	assert.Equal(t, Vert(0.75, 0.25), s.vertices[2])
	assert.Equal(t, Vert(1, 1), s.vertices[3])
}

func TestShape2d_Rotate(t *testing.T) {
	s := Shape(
		Vert(0, 0),
		Vert(0, 1),
		Vert(1, 1),
		Vert(1, 0),
	)

	s.Rotate()

	assert.Equal(t, Vert(1, 0), s.vertices[0])
	assert.Equal(t, Vert(0, 0), s.vertices[1])
	assert.Equal(t, Vert(0, 1), s.vertices[2])
	assert.Equal(t, Vert(1, 1), s.vertices[3])
}
