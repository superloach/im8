package col8_test

import (
	"fmt"
	"image/color"
	"testing"

	"github.com/superloach/im8/col8"
)

var _ = color.Color(col8.Col8(0))

type col8Case struct {
	col        col8.Col8
	r, g, b, a uint32
}

func (c col8Case) test(t *testing.T) {
	cr, cg, cb, ca := c.col.RGBA()

	if cr != c.r {
		t.Errorf("cr == %04x != c.r == %04x", cr, c.r)
	}

	if cg != c.g {
		t.Errorf("cg == %04x != c.g == %04x", cg, c.g)
	}

	if cb != c.b {
		t.Errorf("cb == %04x != c.b == %04x", cb, c.b)
	}

	if ca != c.a {
		t.Errorf("ca == %04x != c.a == %04x", ca, c.a)
	}
}

var col8Cases = []col8Case{
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
	for i, c := range col8Cases {
		t.Run(fmt.Sprint(i), c.test)
	}
}
