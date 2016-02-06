package radius

import (
	"bufio"
	"encoding/binary"
	"io"
)

// A Packet is a RADIUS message
type Packet struct {
	Code       PacketCode
	ID         Identifier
	Attributes []Attribute

	auth Authenticator
}

func (p *Packet) Length() int16 {
	length := 0
	length += 4  // one for code; one for ID; two for length
	length += 16 // 16 for authenticator

	for _, a := range p.Attributes {
		length += int(a.Length)
	}

	return int16(length)
}

func (p *Packet) Write(wx io.Writer) error {

	w := bufio.NewWriter(wx)

	if err := p.Code.Write(w); err != nil {
		return err
	}

	if err := p.ID.Write(w); err != nil {
		return err
	}

	l := p.Length()

	if err := binary.Write(w, binary.BigEndian, uint16(l)); err != nil {
		return err
	}

	hash, err := p.auth.Calculate(p)
	if err != nil {
		return err
	}

	if _, err := w.Write(hash); err != nil {
		return err
	}

	for _, a := range p.Attributes {
		if err := a.Write(w); err != nil {
			return err
		}
	}

	w.Flush()

	return nil
}

func (p *Packet) Read(rx io.Reader) error {
	r := bufio.NewReader(rx)

	if err := p.Code.Read(r); err != nil {
		return err
	}

	if err := p.ID.Read(r); err != nil {
		return err
	}

	var length uint16

	if err := binary.Read(r, binary.BigEndian, &length); err != nil {
		return err
	}

	// just discard hash for now
	var hash []byte = make([]byte, 16)

	if _, err := r.Read(hash); err != nil {
		return err
	}

	// just discard attributes for now
	var buf []byte = make([]byte, length-20)
	if _, err := r.Read(buf); err != nil {
		return err
	}

	return nil
}
