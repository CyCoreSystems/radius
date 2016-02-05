package radius

import (
	"encoding/binary"
	"io"
)

var (
	AccountingStart Attribute = Attribute{AccountingStatusType, 6, fourOctet(1)}
	AccountingStop  Attribute = Attribute{AccountingStatusType, 6, fourOctet(2)}
	InterimUpdate   Attribute = Attribute{AccountingStatusType, 6, fourOctet(3)}
	AccountingOn    Attribute = Attribute{AccountingStatusType, 6, fourOctet(7)}
	AccountingOff   Attribute = Attribute{AccountingStatusType, 6, fourOctet(8)}
)

// An AttributeValue is a value attached to an attribute
type AttributeValue Writer

// An Attribute is a Key-Value pair attached to a request or response
type Attribute struct {
	Type   AttributeType
	Length int8
	Values []Writer
}

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

func fourOctet(i int) []Writer {
	return []Writer{
		WriterFunc(func(w io.Writer) error {
			return binary.Write(w, binary.BigEndian, int32(i))
		}),
	}
}
