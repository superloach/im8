package im8f

import (
	"encoding/binary"
	"fmt"
	"image"
	"io"

	"github.com/superloach/im8"
	"github.com/superloach/im8/col8"
)

// Encoder contains all of the data that needs to be written in im8f format.
type Encoder struct {
	header []byte
	config *image.Config
	img    *im8.Im8
}

// NewEncoder creates an Encoder for any image, using im8.Convert.
func NewEncoder(src image.Image) *Encoder {
	return NewIm8Encoder(im8.Convert(src))
}

// NewIm8Encoder creates an Encoder for an Im8.
func NewIm8Encoder(img *im8.Im8) *Encoder {
	b := img.Bounds()

	return &Encoder{
		header: []byte(Magic),
		config: &image.Config{
			ColorModel: col8.Col8Model,
			Width:      b.Dx(),
			Height:     b.Dy(),
		},
		img: img,
	}
}

func (e *Encoder) writeHeader(w io.Writer) (int, error) {
	if e.header == nil {
		return 0, nil
	}

	n, err := w.Write(e.header)
	if err != nil {
		return n, fmt.Errorf("write header: %w", err)
	}

	e.header = nil

	return n, nil
}

func (e *Encoder) writeConfig(w io.Writer) (int, error) {
	n, err := e.writeHeader(w)
	if err != nil {
		return n, fmt.Errorf("write header: %w", err)
	}

	if e.config == nil {
		return 0, nil
	}

	err = binary.Write(w, binary.BigEndian, uint64(e.config.Width))
	n += 8

	if err != nil {
		return n, fmt.Errorf("write width: %w", err)
	}

	err = binary.Write(w, binary.BigEndian, uint64(e.config.Height))
	n += 8

	if err != nil {
		return n, fmt.Errorf("write height: %w", err)
	}

	e.config = nil

	return n, nil
}

func (e *Encoder) writeImage(w io.Writer) (int, error) {
	n, err := e.writeConfig(w)
	if err != nil {
		return n, fmt.Errorf("write config: %w", err)
	}

	if e.img == nil {
		return 0, nil
	}

	m, err := w.Write(e.img.Pix)
	n += m

	if err != nil {
		return n, fmt.Errorf("write pix: %w", err)
	}

	e.img = nil

	return n, nil
}

// Write writes the contents of the Encoder to w.
func (e *Encoder) Write(w io.Writer) (int, error) {
	if e.img == nil {
		return 0, nil
	}

	return e.writeImage(w)
}

// Encode writes src to w in im8f format.
func Encode(w io.Writer, src image.Image) error {
	_, err := NewEncoder(src).Write(w)
	if err != nil {
		return fmt.Errorf("write to w: %w", err)
	}

	return nil
}
