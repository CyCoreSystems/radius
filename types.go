package radius

import "io"

// AttributeType defines types for an Attribute
type AttributeType int64

const (

	// Attributes for RFC2866/Radius accounting

	// AccountingStatusType indicates whether the AccountingRequest is the beginning of a user service (Start) or the end of a user service (Stop)
	AccountingStatusType AttributeType = 40

	// AccountingDelayTime indicates how many seconds the client has been trying to send this record for
	AccountingDelayTime AttributeType = 41

	// AccountingInputOctets indicates how many octets have been received from this port over the course of this service being provided and is only usable when AccountingStatusType is Stop
	AccountingInputOctets AttributeType = 42

	// AccountingOutputOctets indicates how many octets have been sent from this from this port over the course of this service being provided and is only usable when AccountingStatusType is Stop
	AccountingOutputOctets AttributeType = 43

	// AccountingSessionID indicates the session ID for the accounting session, must match for the matching Start and Stop message, and must be present on all Accounting messages
	AccountingSessionID AttributeType = 44

	// AccountingAuthentic indicates how the user was authenticated and MAY be included.
	AccountingAuthentic AttributeType = 45

	// AccountingSessionTime indicates how many seconds the user has received service for and can only be included when AccountingStatusType is Stop
	AccountingSessionTime AttributeType = 46

	// AccountingInputPackets incidates how many packets have been received from the port over the course of this service being provided to a Framed (?) User and is only usable when AccountingStatusType is Stop
	AccountingInputPackets AttributeType = 47

	// AccountingOutputPackets incidates how many packets have been sent from the port over the course of this service being provided to a Framed (?) User and is only usable when AccountingStatusType is Stop
	AccountingOutputPackets AttributeType = 48

	// AccountingTerminateCause indicates how the session was terminated and is only usable when AccountingStatusType is Stop
	AccountingTerminateCause AttributeType = 49

	// AccountingMultiSessionID is used to mark multiple related sessions together
	AccountingMultiSessionID AttributeType = 50

	// AccountingMultiLinkCount is used to convey how many related sessions are linked together via the AccountingMultiSessionID
	AccountingMultiLinkCount AttributeType = 51
	// --

)

// Write writes the attribute type to the given writer
func (a AttributeType) Write(w io.Writer) error {
	_, err := w.Write([]byte{byte(a)})
	return err
}
