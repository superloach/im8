package col8_test

import (
	"fmt"
	"image/color"
	"testing"

	"github.com/superloach/im8/col8"
)

type col8ModelCase col8Case

func (c col8ModelCase) test(t *testing.T) {
	col8Case{
		col: col8.Col8Model.Convert(color.RGBA64{
			R: uint16(c.r),
			G: uint16(c.g),
			B: uint16(c.b),
			A: uint16(c.a),
		}).(col8.Col8),
		r: c.r,
		g: c.g,
		b: c.b,
		a: c.a,
	}.test(t)
}

func TestCol8Model(t *testing.T) {
	for i, c := range col8Cases {
		t.Run(fmt.Sprint(i), col8ModelCase(c).test)
	}
}
