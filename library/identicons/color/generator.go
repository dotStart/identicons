package color

import "github.com/lucasb-eyer/go-colorful"

// Generator provides colors for icons based on an arbitrary input value as found within a hash.
type Generator interface {

	// RequiredBits returns the minimum amount of bits required to map the entire colorspace provided
	// by this generator.
	RequiredBits() uint

	// Generate returns a color for a given arbitrarily sized code.
	Generate(code uint64) colorful.Color
}
