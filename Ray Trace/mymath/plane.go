package mymath

type Plane struct {
	Origin Vec4
	Normal Vec4
}

func (p Plane) Intersect(r Ray) (Intersect, bool) {
	toPoint := Sub(r.Origin, p.Origin)
	dist := -Dot(toPoint, p.Normal) / Dot(r.Direction, p.Normal)

	if dist <= 0.0 {
		return Intersect{}, false
	}

	colPt := Add(r.Origin, Multiply(r.Direction, dist))
	return Intersect{colPt, p.Normal, p, dist}, true
}
