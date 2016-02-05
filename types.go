package radius

import "io"

// AttributeType defines types for an Attribute
type AttributeType int64

const (

	// Attributes for RFC2866/Radius accounting
	AccountingStatusType     AttributeType = 40
	AccountingDelayTime      AttributeType = 41
	AccountingInputOctets    AttributeType = 42
	AccountingOutputOctets   AttributeType = 43
	AccountingSessionID      AttributeType = 44
	AccountingAuthentic      AttributeType = 45
	AccountingSessionTime    AttributeType = 46
	AccountingInputPackets   AttributeType = 47
	AccountingOutputPackets  AttributeType = 48
	AccountingTerminateCause AttributeType = 49
	AccountingMultiSessionID AttributeType = 50
	AccountingMultiLinkCount AttributeType = 51
	// --

)

// Write writes the attribute type to the given writer
func (a AttributeType) Write(w io.Writer) error {
	_, err := w.Write([]byte{byte(a)})
	return err
}
