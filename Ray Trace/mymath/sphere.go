package mymath

import (
	"math"
)

type Sphere struct {
	Origin Vec4
	Radius float64
}

func (s Sphere) Intersect(ray Ray) (Intersect, bool) {
	q := Sub(s.Origin, ray.Origin)
	c := q.Magnitude()
	v := Dot(q, ray.Direction)
	d := s.Radius*s.Radius - (c*c - v*v)

	if d < 0.0 {
		return Intersect{}, false
	}

	dist := v - math.Sqrt(d)
	if dist < 0.0 {
		return Intersect{}, false
	}

	colPt := Add(ray.Origin, Multiply(ray.Direction, dist))
	colNorm := Unit(Sub(colPt, s.Origin))

	return Intersect{colPt, colNorm, s, dist}, true
}
