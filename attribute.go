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

	UserRequest        = Attribute{AccountingTerminateCause, 6, fourOctet(1)}
	LostCarrier        = Attribute{AccountingTerminateCause, 6, fourOctet(2)}
	LostService        = Attribute{AccountingTerminateCause, 6, fourOctet(3)}
	IdleTimeout        = Attribute{AccountingTerminateCause, 6, fourOctet(4)}
	SessionTimeout     = Attribute{AccountingTerminateCause, 6, fourOctet(5)}
	AdminReset         = Attribute{AccountingTerminateCause, 6, fourOctet(6)}
	AdminReboot        = Attribute{AccountingTerminateCause, 6, fourOctet(7)}
	PortError          = Attribute{AccountingTerminateCause, 6, fourOctet(8)}
	NASError           = Attribute{AccountingTerminateCause, 6, fourOctet(9)}
	NASRequest         = Attribute{AccountingTerminateCause, 6, fourOctet(10)}
	NASReboot          = Attribute{AccountingTerminateCause, 6, fourOctet(11)}
	PortUnneeded       = Attribute{AccountingTerminateCause, 6, fourOctet(12)}
	PortPreempted      = Attribute{AccountingTerminateCause, 6, fourOctet(13)}
	PortSuspended      = Attribute{AccountingTerminateCause, 6, fourOctet(14)}
	ServiceUnavailable = Attribute{AccountingTerminateCause, 6, fourOctet(15)}
	Callback           = Attribute{AccountingTerminateCause, 6, fourOctet(16)}
	UserError          = Attribute{AccountingTerminateCause, 6, fourOctet(17)}
	HostRequest        = Attribute{AccountingTerminateCause, 6, fourOctet(18)}
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
