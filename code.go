package radius

import "io"

// A PacketCode is a code that defines the type of the packet
type PacketCode int64

const (

	// Packet codes for RFC2866/Accounting

	AccountingRequest  PacketCode = 4
	AccountingResponse PacketCode = 5

	// --
)

// Write writes the packet code to the writer
func (code PacketCode) Write(w io.Writer) error {
	_, err := w.Write([]byte{byte(code)})
	return err
}
