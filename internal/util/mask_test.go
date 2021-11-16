package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaskBits64(t *testing.T) {
	assert.Equal(t, uint64(0), MaskBits64(0))

	var expected uint64
	for i := 1; i <= 64; i++ {
		expected = expected | (1 << (i - 1))
		assert.Equal(t, expected, MaskBits64(uint(i)))
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("unreported out of bounds")
		}
	}()

	MaskBits64(65)
}
