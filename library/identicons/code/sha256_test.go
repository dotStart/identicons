package code

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSha256Generator_Generate(t *testing.T) {
	gen := SHA256()

	expected := []byte{
		0x9f, 0x86, 0xd0, 0x81, 0x88, 0x4c, 0x7d, 0x65,
		0x9a, 0x2f, 0xea, 0xa0, 0xc5, 0x5a, 0xd0, 0x15,
		0xa3, 0xbf, 0x4f, 0x1b, 0x2b, 0xb, 0x82, 0x2c,
		0xd1, 0x5d, 0x6c, 0x15, 0xb0, 0xf0, 0xa, 0x8,
	}
	actual := gen.Generate(payload)

	assert.Equal(t, expected, actual)
}

func TestSha256Generator_Generate16(t *testing.T) {
	gen := SHA256()

	expected := uint16(0x9f86)
	actual := gen.Generate16(payload)

	assert.Equal(t, expected, actual)
}

func TestSha256Generator_Generate32(t *testing.T) {
	gen := SHA256()

	expected := uint32(0x9f86d081)
	actual := gen.Generate32(payload)

	assert.Equal(t, expected, actual)
}

func TestSha256Generator_Generate64(t *testing.T) {
	gen := SHA256()

	expected := uint64(0x9f86d081884c7d65)
	actual := gen.Generate64(payload)

	assert.Equal(t, expected, actual)
}
