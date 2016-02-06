package radius

import (
	"bytes"
	"testing"
)
import . "github.com/smartystreets/goconvey/convey"

func TestPacketCodeWriting(t *testing.T) {
	Convey("Given some packet code", t, func() {
		a := AccountingRequest
		w := bytes.NewBuffer([]byte{})

		Convey("When packet code is written", func() {
			err := a.Write(w)
			b := w.Bytes()

			Convey("Error should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Byte Length should be 1", func() {
				So(len(b), ShouldEqual, 1)
			})

			Convey("Bytes should be the single value of the input", func() {
				So(bytes.Equal(b, []byte{byte(a)}), ShouldEqual, true)
			})
		})
	})
}

func TestPacketCodeReading(t *testing.T) {
	Convey("Given an empty packet code", t, func() {
		var a PacketCode
		r := bytes.NewBuffer([]byte{byte(AccountingRequest)})

		Convey("When packet code is read", func() {
			err := a.Read(r)

			Convey("Error should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Packet code should be AccountingRequest", func() {
				So(a, ShouldEqual, AccountingRequest)
			})
		})
	})
}
