package code

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var payload = []byte("test")

func TestSha1Generator_Generate(t *testing.T) {
	gen := SHA1()

	expected := []byte{
		0xa9, 0x4a, 0x8f, 0xe5, 0xcc,
		0xb1, 0x9b, 0xa6, 0x1c, 0x4c,
		0x8, 0x73, 0xd3, 0x91, 0xe9,
		0x87, 0x98, 0x2f, 0xbb, 0xd3,
	}
	actual := gen.Generate(payload)

	assert.Equal(t, expected, actual)
}

func TestSha1Generator_Generate16(t *testing.T) {
	gen := SHA1()

	expected := uint16(0xa94a)
	actual := gen.Generate16(payload)

	assert.Equal(t, expected, actual)
}

func TestSha1Generator_Generate32(t *testing.T) {
	gen := SHA1()

	expected := uint32(0xa94a8fe5)
	actual := gen.Generate32(payload)

	assert.Equal(t, expected, actual)
}

func TestSha1Generator_Generate64(t *testing.T) {
	gen := SHA1()

	expected := uint64(0xa94a8fe5ccb19ba6)
	actual := gen.Generate64(payload)

	assert.Equal(t, expected, actual)
}
