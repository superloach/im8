package im8f

import "fmt"

// Magic is the string that appears at the beginning of every im8f encoding.
const Magic = "\x69IM8\n"

// MagicMismatchError is an error that occurs when the decoded data doesn't match Magic.
type MagicMismatchError struct {
	Index int
	Got   byte
}

// Error returns a string representation of the MagicMismatchError.
func (m MagicMismatchError) Error() string {
	return fmt.Sprintf("magic mismatch at index %d: expected byte %q, got %q", m.Index, []byte(Magic)[m.Index], m.Got)
}
