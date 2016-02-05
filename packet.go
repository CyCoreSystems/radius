package radius

import (
	"bytes"
	"encoding/binary"
	"io"
)

// A Packet is a RADIUS message
type Packet struct {
	Code PacketCode
	ID   Identifier
	//TODO: request authenticator
	Attributes []Attribute
}

func (p *Packet) Write(w io.Writer) error {
	if err := p.Code.Write(w); err != nil {
		return err
	}

	if err := p.ID.Write(w); err != nil {
		return err
	}

	// calculate length

	b := bytes.NewBuffer([]byte{})

	//TODO: request authenticator

	for _, a := range p.Attributes {
		if err := a.Write(b); err != nil {
			return err
		}
	}

	buf := b.Bytes()

	length := 0
	length += 4        // one for code; one for ID; two for length
	length += len(buf) // add buffer

	// --

	padding := bytes.NewBuffer([]byte{})

	// write padding
	for length < 20 {
		length++
		if _, err := padding.Write([]byte{0}); err != nil {
			return err
		}
	}

	// write length
	if err := binary.Write(w, binary.BigEndian, int16(length)); err != nil {
		return err
	}

	// write rest of body
	if _, err := w.Write(buf); err != nil {
		return err
	}

	// write padding
	if _, err := w.Write(padding.Bytes()); err != nil {
		return err
	}

	return nil
}
