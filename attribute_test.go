package radius

import (
	"bytes"
	"testing"
)
import . "github.com/smartystreets/goconvey/convey"

func TestAttribute(t *testing.T) {
	Convey("Given AccountingStop Packet Attribute", t, func() {
		a := AccountingStop
		w := bytes.NewBuffer([]byte{})

		Convey("When attribute is written", func() {
			err := a.Write(w)
			b := w.Bytes()

			Convey("Error should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Byte Length should be attribute.Length", func() {
				So(len(b), ShouldEqual, a.Length)
			})

			Convey("Bytes should be 40, 6, 0, 0, 0, 2", func() {
				v := []byte{40, 6, 0, 0, 0, 2}
				So(bytes.Equal(b, v), ShouldEqual, true)
			})
		})
	})

	Convey("Given AccountingOff Packet Attribute", t, func() {
		a := AccountingOff
		w := bytes.NewBuffer([]byte{})

		Convey("When attribute is written", func() {
			err := a.Write(w)
			b := w.Bytes()

			Convey("Error should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Byte Length should be attribute.Length", func() {
				So(len(b), ShouldEqual, a.Length)
			})

			Convey("Bytes should be 40, 6, 0, 0, 0, 8", func() {
				v := []byte{40, 6, 0, 0, 0, 8}
				So(bytes.Equal(b, v), ShouldEqual, true)
			})
		})
	})

}
