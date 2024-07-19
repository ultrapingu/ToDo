package render

import (
	"image"
	"image/png"
	"os"
)

type Buffer struct {
	Width, Height int
	Img           []FRGBA
}

func CreateBuffer(width int, height int) *Buffer {
	return &Buffer{
		Width:  width,
		Height: height,
		Img:    make([]FRGBA, width*height),
	}
}

func (b *Buffer) SetAt(x int, y int, color FRGBA) {
	b.Img[x+y*int(b.Width)] = color
}

func (b *Buffer) GetAt(x int, y int) FRGBA {
	return b.Img[x+y*int(b.Width)]
}

func (b *Buffer) Clear(clearColor FRGBA) {
	for x := 0; x < b.Width; x++ {
		for y := 0; y < b.Height; y++ {
			b.SetAt(x, y, clearColor)
		}
	}
}

func (b *Buffer) Save(filename string) {
	f, err := os.Create(filename)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	img := image.NewNRGBA(image.Rect(0, 0, b.Width, b.Height))
	for x := 0; x < b.Width; x++ {
		for y := 0; y < b.Height; y++ {
			img.SetNRGBA(x, y, b.Img[x+y*int(b.Width)].ToNRGBA())
		}
	}

	png.Encode(f, img)
}
