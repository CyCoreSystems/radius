package radius

import (
	"encoding/binary"
	"io"
)

var (

	// AccountingStart represents the start of RADIUS accounting session for a specific User
	AccountingStart = Attribute{AccountingStatusType, 6, fourOctet(1)}

	// AccountingStop represents the stop of a RADIUS accounting session for a specific USer
	AccountingStop = Attribute{AccountingStatusType, 6, fourOctet(2)}

	// The spec is NOT helpful on what these are for

	// ?
	InterimUpdate = Attribute{AccountingStatusType, 6, fourOctet(3)}

	// ?
	AccountingOn = Attribute{AccountingStatusType, 6, fourOctet(7)}

	// ?
	AccountingOff = Attribute{AccountingStatusType, 6, fourOctet(8)}
)

// An AttributeValue is a value attached to an attribute
type AttributeValue Writer

// An Attribute is a Key-Value pair attached to a request or response
type Attribute struct {
	Type   AttributeType
	Length int8
	Values []Writer
}

// Write writes the attribute to the given writer
func (attr *Attribute) Write(w io.Writer) error {
	if err := attr.Type.Write(w); err != nil {
		return err
	}

	if err := binary.Write(w, binary.BigEndian, attr.Length); err != nil {
		return err
	}

	for _, v := range attr.Values {
		if err := v.Write(w); err != nil {
			return err
		}
	}

	return nil
}

// shortcut to represent a value as 4 octets
func fourOctet(i int) []Writer {
	return []Writer{
		WriterFunc(func(w io.Writer) error {
			return binary.Write(w, binary.BigEndian, int32(i))
		}),
	}
}
