package im8f

import (
	"encoding/binary"
	"fmt"
	"image"
	"io"

	"github.com/superloach/im8"
	"github.com/superloach/im8/col8"
)

const magic = "\x69IM8\n"

type HeaderMismatchError [2]byte

func (h HeaderMismatchError) Error() string {
	return fmt.Sprintf("expected byte %q, got %q", h[0], h[1])
}

type Decoder struct {
	header []byte
	config *image.Config
	img    *im8.Im8
}

func NewDecoder() *Decoder {
	return &Decoder{}
}

func (d *Decoder) readHeader(r io.Reader) (int, error) {
	n := 0

	ohdr := []byte(magic)

	if d.header == nil {
		hdr := make([]byte, len(ohdr))

		m, err := r.Read(hdr)
		n += m
		if err != nil {
			return n, fmt.Errorf("read hdr: %w", err)
		}

		d.header = hdr
	}

	for i, b := range d.header {
		if ohdr[i] != b {
			return n, HeaderMismatchError{ohdr[i], b}
		}
	}

	return n, nil
}

func (d *Decoder) readConfig(r io.Reader) (int, error) {
	n := 0

	m, err := d.readHeader(r)
	n += m
	if err != nil {
		return n, fmt.Errorf("read header: %w", err)
	}

	if d.config != nil {
		return n, nil
	}

	conf := image.Config{
		ColorModel: col8.Col8Model,
	}

	w := uint64(0)
	err = binary.Read(r, binary.BigEndian, &w)
	n += 8
	if err != nil {
		return n, fmt.Errorf("read w: %w", err)
	}

	conf.Width = int(w)

	h := uint64(0)
	err = binary.Read(r, binary.BigEndian, &h)
	n += 8
	if err != nil {
		return n, fmt.Errorf("read h: %w", err)
	}

	conf.Height = int(h)

	d.config = &conf

	return n, nil
}

func (d *Decoder) Config(r io.Reader) (image.Config, error) {
	if d.config != nil {
		return *d.config, nil
	}

	_, err := d.readConfig(r)
	if err != nil {
		return image.Config{}, fmt.Errorf("read config: %w", err)
	}

	return *d.config, nil
}

func (d *Decoder) readImage(r io.Reader) (int, error) {
	n := 0

	m, err := d.readConfig(r)
	n += m
	if err != nil {
		return n, fmt.Errorf("read config: %w", err)
	}

	if d.img != nil {
		return n, nil
	}

	img := &im8.Im8{}

	img.Pix = make([]uint8, d.config.Width*d.config.Height)
	img.Stride = d.config.Width
	img.Rect = image.Rect(0, 0, d.config.Width, d.config.Height)

	m, err = r.Read(img.Pix)
	n += m
	if err != nil {
		return n, fmt.Errorf("read pix: %w", err)
	}

	d.img = img

	return n, nil
}

func (d *Decoder) Image(r io.Reader) (image.Image, error) {
	if d.img != nil {
		return d.img, nil
	}

	_, err := d.readImage(r)
	if err != nil {
		return nil, fmt.Errorf("read image: %w", err)
	}

	return d.img, nil
}

func Decode(r io.Reader) (image.Image, error) {
	return NewDecoder().Image(r)
}

func DecodeConfig(r io.Reader) (image.Config, error) {
	return NewDecoder().Config(r)
}

func init() {
	image.RegisterFormat("im8f", magic, Decode, DecodeConfig)
}
