package im8

import (
	"image"
	"image/draw"
)

// Convert turns any image.Image into an Im8.
func Convert(src image.Image) *Im8 {
	b := src.Bounds()
	img := NewIm8(image.Rect(0, 0, b.Dx(), b.Dy()))

	draw.Draw(img, img.Bounds(), src, b.Min, draw.Src)

	return img
}
