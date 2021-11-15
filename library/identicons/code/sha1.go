package code

import (
	"crypto/sha1"
	"encoding/binary"
)

type sha1Generator struct {
}

var sha1GeneratorInstance = &sha1Generator{}

// SHA1 returns a generator which relies on sha1.Sum in order to generate a code for a given input
// value.
func SHA1() Generator {
	return sha1GeneratorInstance
}

func (*sha1Generator) Generate(input []byte) []byte {
	sum := sha1.Sum(input)
	return sum[:]
}

func (g *sha1Generator) Generate16(input []byte) uint16 {
	sum := g.Generate(input)
	return binary.BigEndian.Uint16(sum)
}

func (g *sha1Generator) Generate32(input []byte) uint32 {
	sum := g.Generate(input)
	return binary.BigEndian.Uint32(sum)
}

func (g *sha1Generator) Generate64(input []byte) uint64 {
	sum := g.Generate(input)
	return binary.BigEndian.Uint64(sum)
}
