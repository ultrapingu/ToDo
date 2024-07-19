package mymath

type Ray struct {
	Direction Vec4
	Origin    Vec4
}

func NewRay(direction, origin Vec4) Ray {
	return Ray{direction, origin}
}

func (r *Ray) Normalize() {
	r.Direction.Normalize()
}
