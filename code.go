package radius

import "io"

// A PacketCode is a code that defines the type of the packet
type PacketCode int64

const (

	// Packet codes for RFC2866/Accounting

	// AccountingRequest represents a RADIUS accounting request packet
	AccountingRequest PacketCode = 4

	// AccountingResponse respresents a RADIUS accounting response packet
	AccountingResponse PacketCode = 5

	// --
)

// Write writes the packet code to the writer
func (code PacketCode) Write(w io.Writer) error {
	_, err := w.Write([]byte{byte(code)})
	return err
}

// Read reads the packet code from the reader
func (code *PacketCode) Read(r io.Reader) error {
	b := make([]byte, 1)
	_, err := r.Read(b)

	*code = PacketCode(b[0])

	return err
}
