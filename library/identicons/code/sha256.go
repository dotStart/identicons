package code

import (
	"crypto/sha256"
	"encoding/binary"
)

type sha256Generator struct {
}

var sha256GeneratorInstance = &sha256Generator{}

// SHA256 returns a generator which relies on sha256.Sum256 in order to generate a code for a given
// input value.
func SHA256() Generator {
	return sha256GeneratorInstance
}

func (*sha256Generator) Generate(input []byte) []byte {
	sum := sha256.Sum256(input)
	return sum[:]
}

func (g *sha256Generator) Generate16(input []byte) uint16 {
	sum := g.Generate(input)
	return binary.BigEndian.Uint16(sum)
}

func (g *sha256Generator) Generate32(input []byte) uint32 {
	sum := g.Generate(input)
	return binary.BigEndian.Uint32(sum)
}

func (g *sha256Generator) Generate64(input []byte) uint64 {
	sum := g.Generate(input)
	return binary.BigEndian.Uint64(sum)
}
