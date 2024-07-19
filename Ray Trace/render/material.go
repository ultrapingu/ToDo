package render

type Material struct {
	Diffuse      FRGBA
	Specular     FRGBA
	SpecularHard float64
	SpecularPow  float64
	Absorbtion   float64
}
