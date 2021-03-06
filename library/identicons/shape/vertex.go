package shape

import "math"

var zero = Vert(0, 0)

type Vert2d struct {
	X float64
	Y float64
}

func Zero() *Vert2d {
	c := *zero
	return &c
}

func Vert(x float64, y float64) *Vert2d {
	return &Vert2d{
		X: x,
		Y: y,
	}
}

func VertAngle(r float64, angle float64) *Vert2d {
	x := r * math.Cos(angle)
	y := r * math.Sin(angle)

	return Vert(x, y)
}

func (v *Vert2d) Invert() {
	v.X = -v.X
	v.Y = -v.Y
}

func (v *Vert2d) Rotate() {
	x := v.X
	y := v.Y

	v.X = -y
	v.Y = x
}

func (v *Vert2d) Multiply(other *Vert2d) {
	v.X *= other.X
	v.Y *= other.Y
}

func (v *Vert2d) Plus(other *Vert2d) {
	v.X += other.X
	v.Y += other.Y
}

func MinVert(verts ...*Vert2d) *Vert2d {
	if len(verts) == 0 {
		return nil
	}
	if len(verts) == 1 {
		return verts[0]
	}

	x := float64(0)
	y := float64(0)
	for i, v := range verts {
		if i == 0 {
			x = v.X
			y = v.Y
			continue
		}

		x = math.Min(x, v.X)
		y = math.Min(y, v.Y)
	}

	return Vert(x, y)
}
