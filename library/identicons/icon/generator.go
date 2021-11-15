package icon

import (
	svg "github.com/ajstarks/svgo"
	"io"
)

// ContentType describes the content type which should be echoed back to clients when writing
// generated icons to http responses.
const ContentType = "image/svg+xml"

// Generator provides a generation function for a given style of identicon.
type Generator interface {

	// Id provides a simple url friendly identifier via which this generator may be referenced when
	// multiple are exposed.
	Id() string

	// Size identifies the width and height of this particular icon as defined within the generated
	// vector file.
	Size() uint

	// Write produces a new image for the given payload and writes its contents to the given writer.
	Write(input []byte, writer io.Writer)

	// Draw generates an icon for the given payload and writes its contents to the SVG at the current
	// location.
	Draw(input []byte, canvas *svg.SVG)
}
