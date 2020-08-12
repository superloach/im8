package im8f

import (
	"encoding/binary"
	"fmt"
	"image"
	"io"

	"github.com/superloach/im8"
	"github.com/superloach/im8/col8"
)

type Encoder struct {
	header []byte
	config *image.Config
	img    *im8.Im8
}

func NewEncoder(img *im8.Im8) *Encoder {
	bounds := img.Bounds()

	return &Encoder{
		header: []byte(Magic),
		config: &image.Config{
			ColorModel: col8.Col8Model,
			Width:      bounds.Dx(),
			Height:     bounds.Dy(),
		},
		img: img,
	}
}

func (e *Encoder) writeHeader(w io.Writer) (int, error) {
	n := 0

	if e.header == nil {
		return 0, nil
	}

	m, err := w.Write(e.header)
	n += m
	if err != nil {
		return n, fmt.Errorf("write header: %w", err)
	}

	e.header = nil

	return n, nil
}

func (e *Encoder) writeConfig(w io.Writer) (int, error) {
	n := 0

	m, err := e.writeHeader(w)
	n += m
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
	n := 0

	m, err := e.writeConfig(w)
	n += m
	if err != nil {
		return n, fmt.Errorf("write config: %w", err)
	}

	if e.img == nil {
		return 0, nil
	}

	m, err = w.Write(e.img.Pix)
	n += m
	if err != nil {
		return n, fmt.Errorf("write pix: %w", err)
	}

	e.img = nil

	return n, nil
}

func (e *Encoder) Write(w io.Writer) (int, error) {
	return e.writeImage(w)
}

func Encode(w io.Writer, img *im8.Im8) error {
	_, err := NewEncoder(img).Write(w)
	if err != nil {
		return fmt.Errorf("write to w: %w", err)
	}

	return nil
}
