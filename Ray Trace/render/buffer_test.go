package render

import (
	"image"
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBufferReturnsANewBuffer(t *testing.T) {
	scene := CreateBuffer(4, 4)
	rect := image.Rect(0, 0, 4, 4)
	assert.Equal(t, 4, scene.Width, "sets width of the buffer")
	assert.Equal(t, 4, scene.Height, "sets height of the buffer")
	assert.True(t, assert.ObjectsAreEqualValues(rect, scene.Img.Bounds()),
		"creates an image.RGBA with proper bounds")
}

func TestBufferSavePanicsWhenFileCannotBePersisted(t *testing.T) {
	scene := CreateBuffer(4, 4)
	clearColor := color.NRGBA{uint8(0), uint8(255), uint8(0), 255}
	scene.Clear(clearColor)
	assert.Panics(t, func() {
		scene.Save("/etc/temp.png")
	}, "panics when the file cannot be persisted")
}
