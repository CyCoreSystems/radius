package radius

import (
	"bytes"
	"testing"
)
import . "github.com/smartystreets/goconvey/convey"

func TestIDWrite(t *testing.T) {
	Convey("Given an identifier", t, func() {
		a := Identifier(12)
		w := bytes.NewBuffer([]byte{})

		Convey("When identifier is written", func() {
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

func TestIDRead(t *testing.T) {

	Convey("Given an empty identifier", t, func() {
		a := Identifier(0)
		r := bytes.NewBuffer([]byte{12})

		Convey("When identifier is read", func() {
			err := a.Read(r)

			Convey("Error should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Identifier should be 12", func() {
				So(a, ShouldEqual, 12)
			})
		})
	})
}
