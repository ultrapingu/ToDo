package mymath

type Shape interface {
	Intersect(ray Ray) (Intersect, bool)
}
