// Package im8 provides a 6- or 8-bit image format.
package im8

import (
	"image"
	"image/color"

	"github.com/superloach/im8/col8"
)

// Im8 is an in-memory image whose At method returns col8.Col8 values.
type Im8 struct {
	// Pix holds the image's pixels. The pixel at (x, y) is at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)].
	Pix []uint8

	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int

	// Rect is the image's bounds.
	Rect image.Rectangle
}

// NewIm8 returns a new Im8 image with the given bounds.
func NewIm8(r image.Rectangle) *Im8 {
	return &Im8{
		Pix:    make([]uint8, r.Dx()*r.Dy()),
		Stride: r.Dx(),
		Rect:   r,
	}
}

// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
func (i *Im8) At(x, y int) color.Color {
	return i.Col8At(x, y)
}

// Bounds returns the domain for which At can return non-zero color.
// The bounds do not necessarily contain the point (0, 0).
func (i *Im8) Bounds() image.Rectangle {
	return i.Rect
}

// ColorModel returns the Image's color model.
func (i *Im8) ColorModel() color.Model {
	return col8.Col8Model
}

// Col8At returns the col8.Col8 pixel at (x, y).
func (i *Im8) Col8At(x, y int) col8.Col8 {
	if !(image.Point{x, y}.In(i.Rect)) {
		return 0
	}

	idx := i.PixOffset(x, y)

	return col8.Col8(i.Pix[idx])
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (i *Im8) Opaque() bool {
	for x := i.Rect.Min.X; x < i.Rect.Max.X; x++ {
		for y := i.Rect.Min.Y; y < i.Rect.Max.Y; y++ {
			_, _, _, a := i.At(x, y).RGBA()
			if a < (1<<16)-1 {
				return false
			}
		}
	}

	return true
}

// PixOffset returns the index of the element of Pix that corresponds to the pixel at (x, y).
func (i *Im8) PixOffset(x, y int) int {
	return (y-i.Rect.Min.Y)*i.Stride + (x - i.Rect.Min.X)
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (i *Im8) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(i.Rect)
	if r.Empty() {
		return &Im8{}
	}

	idx := i.PixOffset(r.Min.X, r.Min.Y)

	return &Im8{
		Pix:    i.Pix[idx:],
		Stride: i.Stride,
		Rect:   r,
	}
}

// Set sets the pixel at (x, y) to c, after converting to a col8.Col8.
func (i *Im8) Set(x, y int, c color.Color) {
	i.SetCol8(x, y, col8.Col8Model.Convert(c).(col8.Col8))
}

// SetCol8 sets the pixel at (x, y) to c.
func (i *Im8) SetCol8(x, y int, c col8.Col8) {
	if !(image.Point{x, y}.In(i.Rect)) {
		return
	}

	idx := i.PixOffset(x, y)

	i.Pix[idx] = uint8(c)
}
