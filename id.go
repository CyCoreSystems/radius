package radius

import (
	"encoding/binary"
	"io"
)

// Identifier is the unique ID of a Packet
type Identifier uint8

func (id Identifier) Write(w io.Writer) error {
	return binary.Write(w, binary.BigEndian, id)
}
