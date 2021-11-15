package code

// Generator encapsulates a hash function which results in a 32 or 64-bit code.
type Generator interface {

	// Generate generates a code of arbitrary size.
	Generate(input []byte) []byte

	// Generate16 generates a 16-bit code for a given input value.
	Generate16(input []byte) uint16

	// Generate32 generates a 32-bit code for a given input value.
	Generate32(input []byte) uint32

	// Generate64 generates a 64-bit code for a given input value.
	Generate64(input []byte) uint64
}
