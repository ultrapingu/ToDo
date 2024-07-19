package main

import (
	mymath "main/mymath"
	render "main/render"
	"math/rand/v2"

	"errors"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseArgs() (int, int, uint32, string, error) {
	args := os.Args[1:]

	if len(args) < 4 {
		return 0, 0, 0, "", errors.New("invalid number of arguments")
	}

	width, err := strconv.Atoi(args[0])
	if err != nil {
		return 0, 0, 0, "", errors.New("imageWidth is not a valid integer")
	} else if width < 0 || width > 4096 {
		return 0, 0, 0, "", errors.New("imageWidth is not in the valid range of 0 - 4096")
	}

	height, err := strconv.Atoi(args[1])
	if err != nil {
		return 0, 0, 0, "", errors.New("imageHeight is not a valid integer")
	} else if height < 0 || height > 4096 {
		return 0, 0, 0, "", errors.New("imageHeight is not in the valid range of 0 - 4096")
	}

	bounceDepth, err := strconv.Atoi(args[2])
	if err != nil {
		return 0, 0, 0, "", errors.New("bounceDepth is not a valid integer")
	} else if bounceDepth <= 0 {
		return 0, 0, 0, "", errors.New("bounceDepth must be a positive integer")
	}

	outputName := args[3]
	if !strings.HasSuffix(strings.ToLower(outputName), ".png") {
		return 0, 0, 0, "", errors.New("outputName must be a png")
	}

	return width, height, uint32(bounceDepth), outputName, nil
}

func loadScene(bounceDepth uint32) render.Scene {
	scene := render.NewScene(bounceDepth)

	minBound := mymath.NewVector(-20.0, -10.0, 3.0)
	maxBound := mymath.NewVector(20.0, 10.0, 20.0)
	minRad := 0.1
	maxRad := 1.5
	numObjs := 100

	defaultMat := render.Material{
		Diffuse:      render.FRGBA{0.2, 0.4, 0.9, 0.0},
		Specular:     render.FRGBA{1.0, 1.0, 1.0, 0.0},
		SpecularHard: 10.0,
		SpecularPow:  10.0,
		Absorbtion:   0.3,
	}

	scene.Objects = make([]render.RenderObj, numObjs)
	for i := range numObjs {
		color := render.FRGBA{rand.Float64(), rand.Float64(), rand.Float64(), 0.0}

		mat := render.Material{
			Diffuse:      color,
			Specular:     color,
			SpecularHard: 1.0 + rand.Float64()*9.0,
			SpecularPow:  1.0 + rand.Float64()*9.0,
			Absorbtion:   rand.Float64(),
		}

		x := minBound.X + rand.Float64()*(maxBound.X-minBound.X)
		y := minBound.Y + rand.Float64()*(maxBound.Y-minBound.Y)
		z := minBound.Z + rand.Float64()*(maxBound.Z-minBound.Z)
		r := minRad + rand.Float64()*(maxRad-minRad)
		scene.Objects[i] = render.RenderObj{Shape: mymath.Sphere{Origin: mymath.NewPoint(x, y, z), Radius: r}, Material: mat}
	}
	scene.Objects = append(scene.Objects, render.RenderObj{Shape: mymath.Plane{Origin: mymath.NewPoint(0.0, minBound.Y, 0.0), Normal: mymath.NewPoint(0.0, 1.0, 0.0)}, Material: defaultMat})
	//scene.Objects = append(scene.Objects, mymath.Plane{Origin: mymath.NewPoint(0.0, maxBound.Y, 0.0), Normal: mymath.NewPoint(0.0, -1.0, 0.0)})

	// scene.Objects = []mymath.Shape{
	// 	mymath.Sphere{Origin: mymath.NewPoint(-2.0, -2.0, 3.0), Radius: 1.0},
	// 	mymath.Sphere{Origin: mymath.NewPoint(2.0, -2.0, 3.0), Radius: 1.0},
	// 	mymath.Sphere{Origin: mymath.NewPoint(0.0, 0.0, 3.0), Radius: 1.0},
	// 	mymath.Sphere{Origin: mymath.NewPoint(-2.0, 2.0, 3.0), Radius: 1.0},
	// 	mymath.Sphere{Origin: mymath.NewPoint(2.0, 2.0, 3.0), Radius: 1.0},
	// 	mymath.Plane{Origin: mymath.NewPoint(0.0, -5.0, 0.0), Normal: mymath.NewPoint(0.0, 1.0, 0.0)},
	// }

	// scene.Objects = []mymath.Shape{
	// 	mymath.Sphere{Origin: mymath.NewPoint(0.0, 0.0, 3.0), Radius: 1.0},
	// 	mymath.Plane{Origin: mymath.NewPoint(0.0, -1.0, 0.0), Normal: mymath.NewPoint(0.0, 1.0, 0.0)},
	// }

	scene.Lights = []render.Light{
		render.AmbLight{Color: render.FRGBA{0.05, 0.05, 0.05, 1.0}},
		render.DirLight{InvDir: mymath.Unit(mymath.NewVector(0.0, 1.0, 0.0)), Color: render.FRGBA{1.0, 1.0, 1.0, 1.0}},
	}

	scene.Background = render.LoadCubeMap("textures/scene.bmp")
	//scene.Background = render.LoadCubeMap("textures/white.bmp")
	//scene.Background = render.SkyGradient{Colors: []render.FRGBA{render.FRGBA{0.27, 0.45, 0.74, 1.0}, render.FRGBA{0.77, 0.86, 0.93, 1.0}, render.FRGBA{0.27, 0.45, 0.74, 1.0}}, Up: mymath.NewVector(0.0, 1.0, 0.0)}

	return scene
}

type PixelQuery struct {
	x, y  int
	scene *render.Scene
	ray   mymath.Ray
}

type PixelResult struct {
	x, y  int
	color render.FRGBA
}

func renderWorker(jobs <-chan PixelQuery, results chan<- PixelResult) {
	for job := range jobs {
		job.ray.Direction.Normalize()
		col := job.scene.Intersect(job.ray)

		results <- PixelResult{job.x, job.y, col}
	}
}

func renderScene(buffer *render.Buffer, scene render.Scene) {
	buffer.Clear(render.FRGBA{1.0, 0.0, 1.0, 1.0})

	pixelCount := buffer.Width * buffer.Height
	jobs := make(chan PixelQuery, pixelCount)
	results := make(chan PixelResult, pixelCount)

	fov := math.Pi / 2.0
	hTanFov := math.Tan(fov / 2.0)
	fWidth := float64(buffer.Width)
	fHeight := float64(buffer.Height)
	aspectRatio := fWidth / fHeight
	rayOrigin := mymath.NewPoint(0.0, 0.0, 0.0)

	for i := 0; i < 100; i++ {
		go renderWorker(jobs, results)
	}

	for x := 0; x < buffer.Width; x++ {
		for y := 0; y < buffer.Height; y++ {
			fX := float64(x) + 0.5
			fY := float64(y) + 0.5

			pX := (2.0*(fX/fWidth) - 1.0) * hTanFov * aspectRatio
			pY := (1.0 - 2.0*(fY/fHeight)) * hTanFov

			rayDirection := mymath.Sub(mymath.NewPoint(pX, pY, 1.0), rayOrigin)
			ray := mymath.NewRay(rayDirection, rayOrigin)

			jobs <- PixelQuery{x, y, &scene, ray}
		}
	}
	close(jobs)

	for i := 0; i < pixelCount; i++ {
		result := <-results
		buffer.SetAt(result.x, result.y, result.color)
	}
}

func main() {
	w, h, bounceDepth, name, err := parseArgs()
	if err != nil {
		log.Fatalf("Invalid arguments: %s. Usage: %s <imageWidth> <imageHeight> <outputName>\n", err.Error(), os.Args[0])
	}

	buffer := render.CreateBuffer(w, h)
	scene := loadScene(bounceDepth)

	renderScene(buffer, scene)
	buffer.Save(name)
}
