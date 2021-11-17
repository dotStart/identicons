package circle

import (
	"fmt"
	svg "github.com/ajstarks/svgo"
	"github.com/dotstart/identicons/library/identicons/code"
	"github.com/dotstart/identicons/library/identicons/color"
	"github.com/dotstart/identicons/library/identicons/icon"
	"github.com/dotstart/identicons/library/identicons/shape"
	"github.com/lucasb-eyer/go-colorful"
	"io"
	"math"
	"math/bits"
)

type generator struct {
	drawCore  bool
	rings     uint
	ringWidth uint
	segments  uint
	size      uint

	centerPoint  *shape.Vert2d
	segmentAngle float64

	codeGenerator code.Generator

	backgroundColor          *colorful.Color
	foregroundColorGenerator color.Generator
}

func New(opts ...Option) icon.Generator {
	g := &generator{
		drawCore:  false,
		rings:     4,
		ringWidth: 32,
		segments:  8,

		codeGenerator:            code.SHA1(),
		foregroundColorGenerator: color.HSL(0.5, 0.45),
	}

	for _, opt := range opts {
		opt(g)
	}

	g.size = (g.rings + 1) * g.ringWidth * 2
	g.segmentAngle = math.Pi * 2 / float64(g.segments)

	return g
}

func (g *generator) Id() string {
	return "circle"
}

func (g *generator) Size() uint {
	return g.size
}

func (g *generator) Write(input []byte, writer io.Writer) {
	canvas := svg.New(writer)
	canvas.Start(int(g.size), int(g.size))
	{
		g.Draw(input, canvas)
	}
	canvas.End()
}

func (g *generator) Draw(input []byte, canvas *svg.SVG) {
	size := int(g.size)
	center := size / 2

	black := colorful.LinearRgb(0, 0, 0)

	c := g.codeGenerator.Generate64(input)

	foregroundColorA := g.foregroundColorGenerator.Generate(c)

	h, s, v := foregroundColorA.Hsv()
	foregroundColorB := colorful.Hsv(h+180, s, v)

	canvas.Group("class=\"identicon\"")
	{
		if g.backgroundColor != nil {
			canvas.Rect(0, 0, size, size, fmt.Sprintf("fill=\"%s\"", g.backgroundColor.Hex()))
		}

		for r := int(g.rings); r >= 0; r-- {
			if r == 0 && !g.drawCore {
				continue
			}

			outerRadius := int(g.ringWidth) * (r + 1)
			innerRadius := outerRadius - int(g.ringWidth)

			canvas.Mask(fmt.Sprintf("ring_%d", r), 0, 0, size, size)
			{
				canvas.Rect(0, 0, size, size, "fill=\"black\"")
				canvas.Circle(center, center, outerRadius, "fill=\"white\"")

				var positions []*shape.Vert2d
				for s := uint(0); s < g.segments; s++ {
					angle := g.segmentAngle * float64(s)

					if (c>>s)&0x01 == 0x00 {
						if positions == nil {
							continue
						}

						endPoint := shape.VertAngle(float64(center*2), angle)
						positions = append(positions, endPoint, shape.Zero())

						s := shape.Shape(positions...)
						s.Draw(canvas, float64(center), float64(center), 1, 1, &black, nil)

						positions = nil
						continue
					}

					if positions == nil {
						positions = make([]*shape.Vert2d, 0)
					}

					pos := shape.VertAngle(float64(center*2), angle)
					positions = append(positions, pos)
				}

				canvas.Circle(center, center, innerRadius, "fill=\"black\"")
				c = bits.RotateLeft64(c, -int(g.segments/2))
			}
			canvas.MaskEnd()

			ringColor := foregroundColorA
			if r%2 == 1 {
				ringColor = foregroundColorB
			}

			canvas.Rect(0, 0, size, size,
				fmt.Sprintf("class=\"ring_%d\"", r),
				fmt.Sprintf("fill=\"%s\"", ringColor.Hex()),
				fmt.Sprintf("mask=\"url(#ring_%d)\"", r))
		}
	}
	canvas.Gend()
}
