package ch6

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConvey(t *testing.T) {
	Convey("Given a starting integer value", t, func() {
		x := 42
		Convey("When incremented", func() {
			x++
			Convey("The value should be greater by one", func() {
				So(x, ShouldEqual, 43)
			})
			Convey("The value should NOT be what it used to be", func() {
				So(x, ShouldNotEqual, 42)
			})
		})
		Convey("When decremented", func() {
			x--
			Convey("The value should be lesser by one", func() {
				So(x, ShouldEqual, 41)
			})
			Convey("The value should NOT be what it used to be", func() {
				So(x, ShouldNotEqual, 42)
			})
		})
	})
}
