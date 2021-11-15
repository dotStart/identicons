package tiled

import (
	"fmt"
	svg "github.com/ajstarks/svgo"
	"github.com/dotstart/identicons/internal/util"
	"github.com/dotstart/identicons/library/identicons/code"
	"github.com/dotstart/identicons/library/identicons/color"
	"github.com/dotstart/identicons/library/identicons/icon"
	"github.com/lucasb-eyer/go-colorful"
	"io"
	"math/bits"
)

const tileCount = 4

const defaultBackgroundColor = "#ffffff"
const defaultTileSize = 64

type Generator struct {
	id string

	codeGenerator code.Generator
	totalBits     uint

	backgroundColor          *colorful.Color
	foregroundColorGenerator color.Generator
	colorSides               bool

	tileSize                 uint
	rotateSides              bool
	permitAdjacentDuplicates bool

	cornerTiles    []Tile
	cornerTileBits uint
	cornerTileMask uint64
	sideTiles      []Tile
	sideTileBits   uint
	sideTileMask   uint64
	centerTiles    []Tile
	centerTileBits uint
	centerTileMask uint64
}

// New creates a new tiled generator using the given set of base parameters.
//
// This function will typically never be called by end users directly but instead be wrapped in a
// template specific function (such as classic.New).
func New(id string, cornerTiles []Tile, sideTiles []Tile, centerTiles []Tile, opts ...Option) icon.Generator {
	cornerTileBits := uint(bits.Len(uint(len(cornerTiles))))
	sideTileBits := uint(bits.Len(uint(len(sideTiles))))
	centerTileBits := uint(bits.Len(uint(len(sideTiles))))

	cornerTileMask := util.MaskBits64(cornerTileBits)
	sideTileMask := util.MaskBits64(sideTileBits)
	centerTileMask := util.MaskBits64(centerTileBits)

	totalBits := cornerTileBits + sideTileBits + centerTileBits
	if totalBits > 64 {
		panic(fmt.Sprintf("tile bits (%d bits) exceed maximum permitted value of 64", totalBits))
	}

	backgroundColor, err := colorful.Hex(defaultBackgroundColor)
	if err != nil {
		panic(fmt.Errorf("failed to parse default background color: %w", err))
	}

	g := &Generator{
		id,

		code.SHA1(),
		totalBits,

		&backgroundColor,
		color.HSV(0.5, 0.45),
		false,

		defaultTileSize,
		false,
		false,

		cornerTiles,
		cornerTileBits,
		cornerTileMask,
		sideTiles,
		sideTileBits,
		sideTileMask,
		centerTiles,
		centerTileBits,
		centerTileMask,
	}

	for _, opt := range opts {
		opt(g)
	}

	return g
}

func (g *Generator) Id() string {
	return g.id
}

func (g *Generator) Size() uint {
	return g.tileSize * tileCount
}

func (g *Generator) Write(input []byte, w io.Writer) {
	size := g.Size()

	canvas := svg.New(w)
	canvas.Start(int(size), int(size))
	{
		g.Draw(input, canvas)
	}
	canvas.End()
}

func (g *Generator) Draw(input []byte, canvas *svg.SVG) {
	c := g.codeGenerator.Generate64(input)

	cornerValue := int(c & g.centerTileMask)
	sideValue := int((c >> g.cornerTileBits) & g.sideTileMask)
	centerValue := int((c >> (g.cornerTileBits + g.sideTileBits)) & g.centerTileMask)

	c = bits.RotateLeft64(c, -int(g.totalBits))

	foregroundColor := g.foregroundColorGenerator.Generate(c)
	c = bits.RotateLeft64(c, -int(g.foregroundColorGenerator.RequiredBits()))

	var sideColor *colorful.Color
	if g.colorSides {
		h, s, v := foregroundColor.Hsv()
		h += 180

		sc := colorful.Hsv(h, s, v)
		sideColor = &sc
	} else {
		sideColor = nil
	}

	cornerRotation := uint(c & 0x3)
	sideRotation := uint((c >> 2) & 0x3)
	centerRotation := uint((c >> 4) & 0x3)

	cornerOffset := cornerValue % len(g.cornerTiles)
	sideOffset := sideValue % len(g.sideTiles)
	centerOffset := centerValue % len(g.centerTiles)

	// prevent the same tile from appearing in adjacent segments
	if !g.permitAdjacentDuplicates {
		if sideOffset == centerOffset {
			sideOffset += 1
			if sideOffset >= len(g.sideTiles) {
				sideOffset = 0
			}
		}
		if cornerOffset == sideOffset {
			cornerOffset += 1
			if cornerOffset >= len(g.cornerTiles) {
				cornerOffset = 0
			}
		}
	}

	cornerTile := g.cornerTiles[cornerOffset]
	sideTile := g.sideTiles[sideOffset]
	centerTile := g.centerTiles[centerOffset]

	tileSize := float64(g.tileSize)
	canvas.Group(fmt.Sprintf("class=\"identicon\" fill=\"%s\"", foregroundColor.Hex()))
	{
		if g.backgroundColor != nil {
			size := int(g.Size())
			canvas.Rect(0, 0, size, size, fmt.Sprintf("class=\"background\" fill=\"%s\"", g.backgroundColor.Hex()))
		}

		// center patch
		canvas.Group("class=\"center\"")
		{
			centerTile.Draw(canvas, tileSize, tileSize, tileSize, false, false, centerRotation)
			centerTile.Draw(canvas, tileSize*2, tileSize, tileSize, true, false, centerRotation)
			centerTile.Draw(canvas, tileSize*2, tileSize*2, tileSize, true, true, centerRotation)
			centerTile.Draw(canvas, tileSize, tileSize*2, tileSize, false, true, centerRotation)
		}
		canvas.Gend()

		// side patches
		if sideColor != nil {
			canvas.Group(fmt.Sprintf("class=\"side\" fill=\"%s\"", sideColor.Hex()))
		} else {
			canvas.Group("class=\"side\"")
		}
		{
			// top #1
			sideTile.Draw(canvas, tileSize, 0, tileSize, false, true, sideRotation)
			sideRotation++

			// right #1
			sideTile.Draw(canvas, tileSize*3, tileSize, tileSize, true, false, sideRotation)
			sideRotation++

			// bottom #1
			sideTile.Draw(canvas, tileSize, tileSize*3, tileSize, true, true, sideRotation)
			sideRotation++

			// left #1
			sideTile.Draw(canvas, 0, tileSize, tileSize, true, true, sideRotation)
			sideRotation++

			if g.rotateSides {
				sideRotation++
			}

			// top #2
			sideTile.Draw(canvas, tileSize*2, 0, tileSize, true, true, sideRotation)
			sideRotation++

			// right #2
			sideTile.Draw(canvas, tileSize*3, tileSize*2, tileSize, true, true, sideRotation)
			sideRotation++

			// bottom #2
			sideTile.Draw(canvas, tileSize*2, tileSize*3, tileSize, false, true, sideRotation)
			sideRotation++

			// left #2
			sideTile.Draw(canvas, 0, tileSize*2, tileSize, true, false, sideRotation)
		}
		canvas.Gend()

		// corner patches
		canvas.Group("class=\"corner\"")
		{
			// top left
			cornerTile.Draw(canvas, 0, 0, tileSize, false, false, cornerRotation)

			// top right
			cornerTile.Draw(canvas, tileSize*3, 0, tileSize, true, false, cornerRotation)

			// bottom right
			cornerTile.Draw(canvas, tileSize*3, tileSize*3, tileSize, true, true, cornerRotation)

			// bottom left
			cornerTile.Draw(canvas, 0, tileSize*3, tileSize, false, true, cornerRotation)
		}
		canvas.Gend()
	}
	canvas.Gend()
}
