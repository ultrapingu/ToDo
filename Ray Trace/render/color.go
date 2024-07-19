package render

import (
	"image/color"
	"main/mymath"
	"math"
)

type FRGBA struct {
	R, G, B, A float64
}

func Black() FRGBA {
	return FRGBA{0.0, 0.0, 0.0, 1.0}
}

func fromFloat(x, y, z, a float64) FRGBA {
	return FRGBA{x, y, z, a}
}

func FromVec4(v mymath.Vec4) FRGBA {
	u := mymath.Unit(v)
	return FRGBA{u.X*0.5 + 0.5, u.Y*0.5 + 0.5, u.Z*0.5 + 0.5, 1.0}
}

func (f *FRGBA) ToNRGBA() color.NRGBA {
	toU8 := func(x float64) uint8 {
		result := int(math.Max(math.Min(x, 1.0), 0.0) * 255)
		if result < 0 {
			return uint8(0)
		} else if result > 255 {
			return uint8(255)
		}

		return uint8(result)
	}
	return color.NRGBA{toU8(f.R), toU8(f.G), toU8(f.B), toU8(f.A)}
}

func Add(a FRGBA, b FRGBA) FRGBA {
	return FRGBA{a.R + b.R, a.G + b.G, a.B + b.B, 1.0}
}

func Multiply(c FRGBA, v float64) FRGBA {
	return FRGBA{c.R * v, c.G * v, c.B * v, c.A}
}

func Blend(a FRGBA, b FRGBA, d float64) FRGBA {
	return FRGBA{a.R + ((b.R - a.R) * d), a.G + ((b.G - a.G) * d), a.B + ((b.B - a.B) * d), 1.0}
}
