package mymath

import (
	"fmt"
	"math"
)

type Vec4 struct {
	X, Y, Z, W float64
}

func NewVec4(x, y, z, w float64) Vec4 {
	return Vec4{x, y, z, w}
}

func NewVector(x, y, z float64) Vec4 {
	return Vec4{x, y, z, 0.0}
}

func NewPoint(x, y, z float64) Vec4 {
	return Vec4{x, y, z, 1.0}
}

func Origin() Vec4 {
	return NewPoint(0.0, 0.0, 0.0)
}

func Up() Vec4 {
	return NewVector(0.0, 1.0, 0.0)
}

func Right() Vec4 {
	return NewVector(1.0, 0.0, 0.0)
}

func Front() Vec4 {
	return NewVector(0.0, 0.0, 1.0)
}

func Add(a Vec4, b Vec4) Vec4 {
	return Vec4{a.X + b.X, a.Y + b.Y, a.Z + b.Z, a.W}
}

func Sub(a Vec4, b Vec4) Vec4 {
	return Vec4{a.X - b.X, a.Y - b.Y, a.Z - b.Z, a.W}
}

func Multiply(a Vec4, b float64) Vec4 {
	return Vec4{a.X * b, a.Y * b, a.Z * b, a.W}
}

func Dot(a Vec4, b Vec4) float64 {
	return (a.X * b.X) + (a.Y * b.Y) + (a.Z * b.Z)
}

func Cross(a Vec4, b Vec4) Vec4 {
	x := a.Y*b.Z - a.Z*b.Y
	y := a.Z*b.X - a.X*b.Z
	z := a.X*b.Y - a.Y*b.X
	return NewVector(x, y, z)
}

func (v *Vec4) MagnitudeSq() float64 {
	return Dot(*v, *v)
}

func (v *Vec4) Magnitude() float64 {
	return float64(math.Sqrt(float64(v.MagnitudeSq())))
}

func Unit(v Vec4) Vec4 {
	magnitude := v.Magnitude()
	return Vec4{v.X / magnitude, v.Y / magnitude, v.Z / magnitude, v.W}
}

func (v *Vec4) Normalize() {
	magnitude := v.Magnitude()
	v.X /= magnitude
	v.Y /= magnitude
	v.Z /= magnitude
}

func (v *Vec4) String() string {
	return fmt.Sprintf("{X: %f, Y: %f, Z: %f, W: %f}", v.X, v.Y, v.Z, v.W)
}
