package radius

import (
	"bytes"
	"testing"
)
import . "github.com/smartystreets/goconvey/convey"

func TestPacket(t *testing.T) {
	Convey("Given Accounting Stop Request Packet", t, func() {
		a := &Packet{
			Code: AccountingRequest,
			ID:   Identifier(12),
			Attributes: []Attribute{
				AccountingStop,
			},
			auth: AccountingRequestAuthenticator("sharedsecret"),
		}

		w := bytes.NewBuffer([]byte{})

		Convey("When attribute is written", func() {
			err := a.Write(w)
			b := w.Bytes()

			Convey("Error should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Byte Length should be 26", func() {
				So(len(b), ShouldEqual, 26)
			})

			//TODO: finish test
			/*
				v := []byte{4, 12, 0, 26, 40, 6, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
				Convey(fmt.Sprintf("Bytes should be %v", v), func() {
					So(b, ShouldResemble, v)
				})
			*/
		})
	})

}
