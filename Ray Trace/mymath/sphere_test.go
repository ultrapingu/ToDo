package mymath

import (
	"fmt"
	"testing"
)

func compareIntersect(a, b Intersect) bool {
	return approxEqualVec4(a.Position, b.Position) &&
		approxEqualVec4(a.Normal, b.Normal) &&
		a.Shape == b.Shape
}

func TestIntersect(t *testing.T) {
	idSphere := Sphere{Origin(), 0.5}

	var testData = []struct {
		sphere    Sphere
		arg1      Ray
		shouldHit bool
		expected  Intersect
	}{
		{idSphere, Ray{Vec4{0.0, 1.0, 0.0, 0.0}, Vec4{0.0, -1.0, 0.0, 1.0}}, true, Intersect{Vec4{0.0, -0.5, 0.0, 1.0}, Vec4{0.0, -1.0, 0.0, 1.0}, idSphere, 0.5}},
		{idSphere, Ray{Vec4{0.0, 1.0, 0.0, 0.0}, Vec4{0.0, 1.0, 0.0, 1.0}}, false, Intersect{}},
		{idSphere, Ray{Vec4{0.0, -1.0, 0.0, 0.0}, Vec4{0.0, 1.0, 0.0, 1.0}}, true, Intersect{Vec4{0.0, 0.5, 0.0, 1.0}, Vec4{0.0, 1.0, 0.0, 1.0}, idSphere, 0.5}},
		{idSphere, Ray{Vec4{0.0, -1.0, 0.0, 0.0}, Vec4{0.0, -1.0, 0.0, 1.0}}, false, Intersect{}},
		{idSphere, Ray{Vec4{1.0, 0.0, 0.0, 0.0}, Vec4{-1.0, 0.0, 0.0, 1.0}}, true, Intersect{Vec4{-0.5, 0.0, 0.0, 1.0}, Vec4{-1.0, 0.0, 0.0, 1.0}, idSphere, 0.5}},
		{idSphere, Ray{Vec4{1.0, 0.0, 0.0, 0.0}, Vec4{1.0, 0.0, 0.0, 1.0}}, false, Intersect{}},
		{idSphere, Ray{Vec4{-1.0, 0.0, 0.0, 0.0}, Vec4{1.0, 0.0, 0.0, 1.0}}, true, Intersect{Vec4{0.5, 0.0, 0.0, 1.0}, Vec4{1.0, 0.0, 0.0, 1.0}, idSphere, 0.5}},
		{idSphere, Ray{Vec4{-1.0, 0.0, 0.0, 0.0}, Vec4{-1.0, 0.0, 0.0, 1.0}}, false, Intersect{}},
		{idSphere, Ray{Vec4{0.0, 0.0, 1.0, 0.0}, Vec4{0.0, 0.0, -1.0, 1.0}}, true, Intersect{Vec4{0.0, 0.0, -0.5, 1.0}, Vec4{0.0, 0.0, -1.0, 1.0}, idSphere, 0.5}},
		{idSphere, Ray{Vec4{0.0, 0.0, 1.0, 0.0}, Vec4{0.0, 0.0, 1.0, 1.0}}, false, Intersect{}},
		{idSphere, Ray{Vec4{0.0, 0.0, -1.0, 0.0}, Vec4{0.0, 0.0, 1.0, 1.0}}, true, Intersect{Vec4{0.0, 0.0, 0.5, 1.0}, Vec4{0.0, 0.0, 1.0, 1.0}, idSphere, 0.5}},
		{idSphere, Ray{Vec4{0.0, 0.0, -1.0, 0.0}, Vec4{0.0, 0.0, -1.0, 1.0}}, false, Intersect{}},
	}

	for _, data := range testData {
		testName := fmt.Sprintf("Test Intersect(%#v, %#v) == %t", data.sphere, data.arg1, data.shouldHit)
		t.Run(testName, func(t *testing.T) {
			intersect, hit := data.sphere.Intersect(data.arg1)

			if hit != data.shouldHit {
				t.Errorf("Intersection results was (%t) but did not match the expected %t\n", hit, data.shouldHit)
				t.FailNow()
			}

			if !compareIntersect(intersect, data.expected) {
				t.Errorf("Intersection result %#v' not equal to expected %#v", intersect, data.expected)
				t.FailNow()
			}
		})
	}
}
