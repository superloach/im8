package col8

import (
	"image/color"
)

// Col8Model is the color.Model for the Col8 type.
var Col8Model color.Model = color.ModelFunc(col8Model)

func col8Model(c color.Color) color.Color {
	if _, ok := c.(Col8); ok {
		return c
	}

	r, g, b, a := c.RGBA()

	r /= channelScale
	g /= channelScale
	b /= channelScale
	a /= channelScale

	a = channelMask - a

	return Col8(a<<6 + r<<4 + g<<2 + b)
}
