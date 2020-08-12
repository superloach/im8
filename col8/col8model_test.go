package col8_test

import (
	"fmt"
	"image/color"
	"testing"

	"github.com/superloach/im8/col8"
)

type col8ModelCase struct {
	r, g, b, a uint16
	col        col8.Col8
}

func (c col8ModelCase) test(t *testing.T) {
	col := col8.Col8Model.Convert(color.RGBA64{
		R: c.r, G: c.g, B: c.b, A: c.a,
	}).(col8.Col8)

	if col != c.col {
		t.Errorf("col == %08b != c.col == %08b", col, c.col)
	}
}

var col8ModelCases = []col8ModelCase{
	{0xFFFF, 0x0000, 0x0000, 0xFFFF, 0b00110000},
	{0xAAAA, 0x0000, 0x0000, 0xFFFF, 0b00100000},
	{0x5555, 0x0000, 0x0000, 0xFFFF, 0b00010000},
	{0x0000, 0x0000, 0x0000, 0xFFFF, 0b00000000},

	{0xFFFF, 0x0000, 0x0000, 0xAAAA, 0b01110000},
	{0xAAAA, 0x0000, 0x0000, 0xAAAA, 0b01100000},
	{0x5555, 0x0000, 0x0000, 0xAAAA, 0b01010000},
	{0x0000, 0x0000, 0x0000, 0xAAAA, 0b01000000},

	{0xFFFF, 0x0000, 0x0000, 0x5555, 0b10110000},
	{0xAAAA, 0x0000, 0x0000, 0x5555, 0b10100000},
	{0x5555, 0x0000, 0x0000, 0x5555, 0b10010000},
	{0x0000, 0x0000, 0x0000, 0x5555, 0b10000000},

	{0xFFFF, 0x0000, 0x0000, 0x0000, 0b11110000},
	{0xAAAA, 0x0000, 0x0000, 0x0000, 0b11100000},
	{0x5555, 0x0000, 0x0000, 0x0000, 0b11010000},
	{0x0000, 0x0000, 0x0000, 0x0000, 0b11000000},

	// should probably write tests for other channels or mixed channels later, idk
}

func TestCol8Model(t *testing.T) {
	for i, c := range col8ModelCases {
		t.Run(fmt.Sprint(i), c.test)
	}
}
