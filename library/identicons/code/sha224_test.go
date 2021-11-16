package code

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSha224Generator_Generate(t *testing.T) {
	gen := SHA224()

	expected := []byte{
		0x90, 0xa3, 0xed, 0x9e,
		0x32, 0xb2, 0xaa, 0xf4,
		0xc6, 0x1c, 0x41, 0xe,
		0xb9, 0x25, 0x42, 0x61,
		0x19, 0xe1, 0xa9, 0xdc,
		0x53, 0xd4, 0x28, 0x6a,
		0xde, 0x99, 0xa8, 0x9,
	}
	actual := gen.Generate(payload)

	assert.Equal(t, expected, actual)
}

func TestSha224Generator_Generate16(t *testing.T) {
	gen := SHA224()

	expected := uint16(0x90a3)
	actual := gen.Generate16(payload)

	assert.Equal(t, expected, actual)
}

func TestSha224Generator_Generate32(t *testing.T) {
	gen := SHA224()

	expected := uint32(0x90a3ed9e)
	actual := gen.Generate32(payload)

	assert.Equal(t, expected, actual)
}

func TestSha224Generator_Generate64(t *testing.T) {
	gen := SHA224()

	expected := uint64(0x90a3ed9e32b2aaf4)
	actual := gen.Generate64(payload)

	assert.Equal(t, expected, actual)
}
