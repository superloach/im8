package col8_test

import (
	"image/color"
	"testing"

	"github.com/superloach/im8/col8"
)

var _ = color.Color(col8.Col8(0))

type Col8Case struct {
	Col        col8.Col8
	R, G, B, A uint32
}

func (c Col8Case) Test(t *testing.T) {
	cr, cg, cb, ca := c.Col.RGBA()

	if cr != c.R {
		t.Errorf("cr == %X != c.R == %X", cr, c.R)
	}

	if cg != c.G {
		t.Errorf("cg == %X != c.G == %X", cg, c.G)
	}

	if cb != c.B {
		t.Errorf("cb == %X != c.B == %X", cb, c.B)
	}

	if ca != c.A {
		t.Errorf("ca == %X != c.A == %X", ca, c.A)
	}
}

var cases = []Col8Case{
	{0b110000, 0xFFFF, 0x0000, 0x0000, 0xFFFF},
	{0b00110000, 0xFFFF, 0x0000, 0x0000, 0xFFFF},
	{0b01110000, 0xAAAA, 0x0000, 0x0000, 0xAAAA},
	{0b10110000, 0x5555, 0x0000, 0x0000, 0x5555},
	{0b11110000, 0x0000, 0x0000, 0x0000, 0x0000},

	{0b100000, 0xAAAA, 0x0000, 0x0000, 0xFFFF},
	{0b00100000, 0xAAAA, 0x0000, 0x0000, 0xFFFF},
	{0b01100000, 0x71C6, 0x0000, 0x0000, 0xAAAA},
	{0b10100000, 0x38E3, 0x0000, 0x0000, 0x5555},
	{0b11100000, 0x0000, 0x0000, 0x0000, 0x0000},

	{0b010000, 0x5555, 0x0000, 0x0000, 0xFFFF},
	{0b00010000, 0x5555, 0x0000, 0x0000, 0xFFFF},
	{0b01010000, 0x38E3, 0x0000, 0x0000, 0xAAAA},
	{0b10010000, 0x1C71, 0x0000, 0x0000, 0x5555},
	{0b11010000, 0x0000, 0x0000, 0x0000, 0x0000},

	{0b000000, 0x0000, 0x0000, 0x0000, 0xFFFF},
	{0b00000000, 0x0000, 0x0000, 0x0000, 0xFFFF},
	{0b01000000, 0x0000, 0x0000, 0x0000, 0xAAAA},
	{0b10000000, 0x0000, 0x0000, 0x0000, 0x5555},
	{0b11000000, 0x0000, 0x0000, 0x0000, 0x0000},

	// should probably write tests for other channels or mixed channels later, idk
}

func TestCol8(t *testing.T) {
	for _, col8Case := range cases {
		col8Case.Test(t)
	}
}
