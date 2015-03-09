package soundex

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSoundex(t *testing.T) {
	Convey("test robert/rupert soundexes", t, func() {
		r1, _ := Soundex("Robert")
		r2, _ := Soundex("Rupert")
		So(r1, ShouldEqual, r2)
	})
	Convey("test Pfister soundexes", t, func() {
		r1, _ := Soundex("Pfister")
		So(r1, ShouldEqual, "P236")
	})
	Convey("test husobee soundexes", t, func() {
		r1, _ := Soundex("husobee")
		So(r1, ShouldEqual, "H210")
	})
	Convey("test Tymczak soundexes", t, func() {
		r1, _ := Soundex("Tymczak")
		So(r1, ShouldEqual, "T522")
	})
	Convey("test Ashcraft soundexes", t, func() {
		r1, _ := Soundex("Ashcraft")
		So(r1, ShouldEqual, "A261")
	})
}
