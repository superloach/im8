// Package col8 provides an 8-bit color type Col8.
package col8

// Col8 is an 8-bit color type, with 2-bit channels in the layout aaRRGGBB.
// Alpha values are inverted, to allow their omission to result in full opacity.
type Col8 uint8

const (
	channelMask  = (1 << 2) - 1
	channelScale = ((1 << 16) - 1) / channelMask
)

// RGBA returns the alpha-premultiplied red, green, blue and alpha values
// for the color. Each value ranges within [0, 0xffff], but is represented
// by a uint32 so that multiplying by a blend factor up to 0xffff will not
// overflow.
//
// An alpha-premultiplied color component c has been scaled by alpha (a),
// so has valid values 0 <= c <= a.
func (c Col8) RGBA() (r, g, b, a uint32) {
	a = (channelMask - uint32(c>>6&channelMask)) * channelScale
	r = (uint32(c>>4&channelMask) * a) / channelMask
	g = (uint32(c>>2&channelMask) * a) / channelMask
	b = (uint32(c&channelMask) * a) / channelMask

	return
}
