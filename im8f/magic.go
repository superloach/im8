package im8f

import "fmt"

const Magic = "\x69IM8\n"

type MagicMismatchError struct {
	Index int
	Got   byte
}

func (m MagicMismatchError) Error() string {
	return fmt.Sprintf("magic mismatch at index %d: expected byte %q, got %q", m.Index, []byte(Magic)[m.Index], m.Got)
}
