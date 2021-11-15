package shape

import (
	"fmt"
	svg "github.com/ajstarks/svgo"
	"github.com/lucasb-eyer/go-colorful"
)

type Shape2d struct {
	vertices []*Vert2d
}

func New(vertices ...*Vert2d) *Shape2d {
	return &Shape2d{vertices}
}

func (s *Shape2d) Flip(flipX bool, flipY bool) {
	if !flipX && !flipY {
		return
	}

	vertices := make([]*Vert2d, len(s.vertices))
	for i, v := range s.vertices {
		vertex := *v

		if flipX {
			vertex.X = 1 - vertex.X
		}
		if flipY {
			vertex.Y = 1 - vertex.Y
		}

		vertices[i] = &vertex
	}

	s.vertices = vertices
}

func (s *Shape2d) Rotate() {
	vertices := make([]*Vert2d, len(s.vertices))
	for i, v := range s.vertices {
		vertex := *v

		vertex.X -= 0.5
		vertex.Y -= 0.5

		vertex.Rotate()

		vertex.X += 0.5
		vertex.Y += 0.5

		vertices[i] = &vertex
	}

	s.vertices = vertices
}

func (s Shape2d) Draw(canvas *svg.SVG, offX float64, offY float64, scaleX float64, scaleY float64, fill *colorful.Color, stroke *colorful.Color) {
	if len(s.vertices) < 2 {
		return
	}

	first := s.vertices[0]
	path := fmt.Sprintf("M%f %f", (first.X*scaleX)+offX, (first.Y*scaleY)+offY)

	for i := 1; i < len(s.vertices); i++ {
		current := s.vertices[i]
		path += fmt.Sprintf(" L%f %f", (current.X*scaleX)+offX, (current.Y*scaleY)+offY)
	}

	attributes := ""
	if fill != nil {
		attributes = fmt.Sprintf("fill=\"%s\"", fill.Hex())
	}
	if stroke != nil {
		if attributes != "" {
			attributes += " "
		}

		attributes += fmt.Sprintf("stroke=\"%s\"", stroke.Hex())
	}

	canvas.Path(path, attributes)
}
