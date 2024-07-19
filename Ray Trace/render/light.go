package render

import (
	"main/mymath"
	"math"
)

type Light interface {
	Light(intersect mymath.Intersect, fromCam mymath.Vec4, mat Material) (FRGBA, FRGBA)
	CastsShadows() bool
	GetShadowRay(i mymath.Intersect) mymath.Ray
}

type AmbLight struct {
	Color FRGBA
}

func (d AmbLight) Light(_ mymath.Intersect, fromCam mymath.Vec4, mat Material) (FRGBA, FRGBA) {
	return FRGBA{d.Color.R * mat.Diffuse.R, d.Color.G * mat.Diffuse.G, d.Color.B * mat.Diffuse.B, 1.0}, Black()
}

func (d AmbLight) CastsShadows() bool {
	return false
}

func (d AmbLight) GetShadowRay(_ mymath.Intersect) mymath.Ray {
	return mymath.Ray{}
}

type DirLight struct {
	InvDir mymath.Vec4
	Color  FRGBA
}

func (d DirLight) Light(intersect mymath.Intersect, fromCam mymath.Vec4, mat Material) (FRGBA, FRGBA) {
	diffuseLight := math.Max(mymath.Dot(intersect.Normal, d.InvDir), 0.0)
	diffuseLightCol := Multiply(d.Color, diffuseLight)
	diffuseCol := FRGBA{diffuseLightCol.R * mat.Diffuse.R, diffuseLightCol.G * mat.Diffuse.G, diffuseLightCol.B * mat.Diffuse.B, 1.0}

	half := mymath.Unit(mymath.Add(fromCam, d.GetShadowRay(intersect).Direction))
	specLight := math.Max(mymath.Dot(intersect.Normal, half), 0.0)
	specLight = math.Pow(specLight, mat.SpecularPow)
	specularLighting := Multiply(d.Color, specLight*mat.SpecularHard)

	specularLighting = FRGBA{specLight * mat.SpecularHard, specLight * mat.SpecularHard, specLight * mat.SpecularHard, 1.0}

	return diffuseCol, specularLighting
}

func (d DirLight) CastsShadows() bool {
	return true
}

func (d DirLight) GetShadowRay(i mymath.Intersect) mymath.Ray {
	return mymath.NewRay(d.InvDir, i.Position)
}
