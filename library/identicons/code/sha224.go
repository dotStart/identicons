package code

import (
	"crypto/sha256"
	"encoding/binary"
)

type sha224Generator struct {
}

var sha224GeneratorInstance = &sha224Generator{}

// SHA224 returns a generator which relies on sha256.Sum224 in order to generate a code for a given
// input value.
func SHA224() Generator {
	return sha224GeneratorInstance
}

func (*sha224Generator) Generate(input []byte) []byte {
	sum := sha256.Sum224(input)
	return sum[:]
}

func (g *sha224Generator) Generate16(input []byte) uint16 {
	sum := g.Generate(input)
	return binary.BigEndian.Uint16(sum)
}

func (g *sha224Generator) Generate32(input []byte) uint32 {
	sum := g.Generate(input)
	return binary.BigEndian.Uint32(sum)
}

func (g *sha224Generator) Generate64(input []byte) uint64 {
	sum := g.Generate(input)
	return binary.BigEndian.Uint64(sum)
}
