package im8f

import (
	"encoding/binary"
	"fmt"
	"image"
	"io"

	"github.com/superloach/im8"
	"github.com/superloach/im8/col8"
)

// Decoder holds the data that is read in im8f format.
type Decoder struct {
	header []byte
	config *image.Config
	img    *im8.Im8
}

// NewDecoder returns an empty Decoder, ready for use.
// &Decoder{} also works fine for this.
func NewDecoder() *Decoder {
	return &Decoder{}
}

func (d *Decoder) readHeader(r io.Reader) (int, error) {
	n := 0
	ohdr := []byte(Magic)

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
			return n, MagicMismatchError{i, b}
		}
	}

	return n, nil
}

func (d *Decoder) readConfig(r io.Reader) (int, error) {
	n, err := d.readHeader(r)
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

// Config reads an image.Config from r in im8f format.
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
	n, err := d.readConfig(r)
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

	m, err := r.Read(img.Pix)
	n += m

	if err != nil {
		return n, fmt.Errorf("read pix: %w", err)
	}

	d.img = img

	return n, nil
}

// Image reads an image.Image from r in im8f format.
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

// Decode reads an image.Image from r in im8f format.
func Decode(r io.Reader) (image.Image, error) {
	return NewDecoder().Image(r)
}

// DecodeConfig reads an image.Config from r in im8f format.
func DecodeConfig(r io.Reader) (image.Config, error) {
	return NewDecoder().Config(r)
}

func init() {
	image.RegisterFormat("im8f", Magic, Decode, DecodeConfig)
}
