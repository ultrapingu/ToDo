package render

import (
	"log"
	"main/mymath"
	"math"
)

type RenderObj struct {
	Shape    mymath.Shape
	Material Material
}

type Scene struct {
	Objects     []RenderObj
	Lights      []Light
	Background  Background
	BounceDepth uint32
}

func NewScene(bounceDepth uint32) Scene {
	if bounceDepth == 0 {
		log.Fatal("bounceDepth must be non zero")
	}

	return Scene{make([]RenderObj, 0), make([]Light, 0), CubeMap{}, bounceDepth}
}

func (s *Scene) GetBackgroundColor(ray mymath.Ray) FRGBA {
	return s.Background.Sample(ray.Direction)
}

func (s *Scene) light(i mymath.Intersect, ignoreIdx int, fromCam mymath.Vec4, mat Material) (FRGBA, FRGBA) {
	diffuse := Black()
	specular := Black()
	for _, l := range s.Lights {
		if l.CastsShadows() {
			shadowRay := l.GetShadowRay(i)
			if _, collisionIdx := s.findNearestIntersectIdx(shadowRay, ignoreIdx); collisionIdx == -1 {
				currD, currS := l.Light(i, fromCam, mat)
				diffuse = Add(diffuse, currD)
				specular = Add(specular, currS)
			}
		} else {
			currD, currS := l.Light(i, fromCam, mat)
			diffuse = Add(diffuse, currD)
			specular = Add(specular, currS)
		}
	}

	return diffuse, specular
}

func (s *Scene) findNearestIntersectIdx(ray mymath.Ray, ignoreIdx int) (mymath.Intersect, int) {
	nearest := math.MaxFloat64
	collidedIdx := -1
	intersect := mymath.Intersect{}

	for idx, obj := range s.Objects {
		if idx != ignoreIdx {
			if inter, hit := obj.Shape.Intersect(ray); hit && inter.Distance < nearest {
				nearest = inter.Distance
				collidedIdx = idx
				intersect = inter
			}
		}
	}

	return intersect, collidedIdx
}

func getReflectedRay(orig mymath.Ray, intersect mymath.Intersect) mymath.Ray {
	td := mymath.Dot(orig.Direction, intersect.Normal) * 2.0
	dpn := mymath.Multiply(intersect.Normal, td)
	direction := mymath.Sub(orig.Direction, dpn)

	return mymath.NewRay(direction, intersect.Position)
}

func (s *Scene) recurse(ray mymath.Ray, cameraPos mymath.Vec4, ignoreIdx int, depth uint32) FRGBA {
	if intersect, collisionIdx := s.findNearestIntersectIdx(ray, ignoreIdx); collisionIdx != -1 {
		mat := s.Objects[collisionIdx].Material

		fromCam := mymath.Unit(mymath.Sub(intersect.Position, cameraPos))

		diffCol, specCol := s.light(intersect, ignoreIdx, fromCam, mat)

		reflectCol := Black()
		if depth != 0 && mat.Absorbtion < 1.0 {
			reflectedRay := getReflectedRay(ray, intersect)
			reflectCol = s.recurse(reflectedRay, cameraPos, collisionIdx, depth-1)
		}

		result := Blend(reflectCol, diffCol, mat.Absorbtion)
		result = Add(result, specCol)

		return result
	}

	return s.GetBackgroundColor(ray)
}

func (s *Scene) Intersect(ray mymath.Ray) FRGBA {
	return s.recurse(ray, ray.Origin, -1, uint32(s.BounceDepth))
}
