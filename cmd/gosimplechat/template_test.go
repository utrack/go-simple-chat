package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestTemplate(t *testing.T) {
	Convey("With template", t, func() {
		t, err := getTemplate()
		Convey("Shouldn't return error", func() {
			So(t, ShouldNotBeNil)
			So(err, ShouldBeNil)
		})
	})
}
