package radius

import (
	"encoding/binary"
	"io"
)

// Identifier is the unique ID of a Packet
type Identifier uint8

// Write writes the identifier to the given writer
func (id Identifier) Write(w io.Writer) error {
	return binary.Write(w, binary.BigEndian, id)
}

// Read reads the identifier from the given reader
func (id *Identifier) Read(r io.Reader) error {
	b := make([]byte, 1)
	_, err := r.Read(b)

	*id = Identifier(b[0])

	return err
}
