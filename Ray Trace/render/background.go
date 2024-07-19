package render

import (
	"fmt"
	"image"
	"log"
	"main/mymath"
	"math"
	"os"
	"strings"

	"golang.org/x/image/bmp"
)

type Background interface {
	Sample(dir mymath.Vec4) FRGBA
}

type CubeMap struct {
	Sides []image.Image
}

const frontIdx = 0
const backIdx = 1
const leftIdx = 2
const rightIdx = 3
const topIdx = 4
const bottomIdx = 5

func LoadCubeMap(filename string) CubeMap {
	splitFilename := strings.Split(filename, ".")
	if len(splitFilename) != 2 {
		log.Fatal("invalid filename")
	}

	result := CubeMap{make([]image.Image, 6)}
	for i, app := range []string{"front", "back", "left", "right", "top", "bottom"} {
		currFn := fmt.Sprintf("%s_%s.%s", splitFilename[0], app, splitFilename[1])

		f, err := os.Open(currFn)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer f.Close()

		img, err := bmp.Decode(f)
		if err != nil {
			log.Fatal(err.Error())
		}

		result.Sides[i] = img
	}

	return result
}

func sampleImage(i image.Image, x float64, y float64) FRGBA {
	x = math.Mod(x, 1.0)
	y = math.Mod(y, 1.0)

	iX := int(x * float64(i.Bounds().Dx()))
	iY := int(y * float64(i.Bounds().Dy()))

	r, g, b, a := i.At(iX, iY).RGBA()
	return FRGBA{float64(r) / float64(a), float64(g) / float64(a), float64(b) / float64(a), 1.0}
}

func (c CubeMap) Sample(dir mymath.Vec4) FRGBA {
	aX := math.Abs(dir.X)
	aY := math.Abs(dir.Y)
	aZ := math.Abs(dir.Z)

	if aX >= aY && aX >= aZ {
		// Right face
		if dir.X > 0.0 {
			x := 1.0 - (dir.Z/dir.X+1.0)*0.5
			y := 1.0 - (dir.Y/dir.X+1.0)*0.5
			return sampleImage(c.Sides[rightIdx], x, y)
		} else {
			x := 1.0 - (dir.Z/dir.X+1.0)*0.5
			y := (dir.Y/dir.X + 1.0) * 0.5
			return sampleImage(c.Sides[leftIdx], x, y)
		}
	} else if aY >= aX && aY >= aZ {
		// Top face
		if dir.Y > 0.0 {
			x := (dir.X/dir.Y + 1.0) * 0.5
			y := (dir.Z/dir.Y + 1.0) * 0.5
			return sampleImage(c.Sides[topIdx], x, y)
		} else {
			x := 1.0 - (dir.X/dir.Y+1.0)*0.5
			y := 1.0 - (dir.Z/dir.Y+1.0)*0.5
			return sampleImage(c.Sides[bottomIdx], x, y)
		}
	} else if aZ >= aX && aZ >= aY {
		// Front face
		if dir.Z > 0.0 {
			x := (dir.X/dir.Z + 1.0) * 0.5
			y := 1.0 - (dir.Y/dir.Z+1.0)*0.5
			return sampleImage(c.Sides[frontIdx], x, y)
		} else {
			x := (dir.X/dir.Z + 1.0) * 0.5
			y := (dir.Y/dir.Z + 1.0) * 0.5
			return sampleImage(c.Sides[backIdx], x, y)
		}
	}

	return Black()
}

type SkyGradient struct {
	Colors []FRGBA
	Up     mymath.Vec4
}

func (s SkyGradient) Sample(dir mymath.Vec4) FRGBA {
	dot := (mymath.Dot(s.Up, dir) * 0.5) + 0.5
	pos := dot * float64(len(s.Colors)-1)
	lower := math.Floor(pos)
	upper := math.Ceil(pos)
	return Blend(s.Colors[int(lower)], s.Colors[int(upper)], pos-lower)
}
