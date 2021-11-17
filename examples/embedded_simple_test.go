package examples

import (
	"bytes"
	"github.com/dotstart/identicons/library/identicons"
	"github.com/dotstart/identicons/library/identicons/code"
	"github.com/dotstart/identicons/library/identicons/color"
	"github.com/dotstart/identicons/library/identicons/icon/modern"
	"github.com/dotstart/identicons/library/identicons/icon/tiled"
	"github.com/lucasb-eyer/go-colorful"
)

// basic usage example for interacting with the generators
func Example_simple() {
	writer := bytes.NewBuffer(make([]byte, 0))

	gen := modern.New(
		tiled.Code(code.SHA256()),
		tiled.BackgroundColor(colorful.LinearRgb(0.95, 0.95, 0.95)),
		tiled.ForegroundColor(color.HSV(0.5, 0.45)),
		tiled.ColoredSides(true),
		tiled.PermitAdjacentDuplicates(true),
	)

	gen.Write([]byte("some string"), writer)
}

func Example_registry() {
	writer := bytes.NewBuffer(make([]byte, 0))
	registry := identicons.DefaultRegistry()

	gen := registry.Get("modern") // from request or config file
	gen.Write([]byte("some string"), writer)
}
