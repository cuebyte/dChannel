package dChannel

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNew(t *testing.T) {
	Convey("Given ", t, func() {
		New()
	})
}

func TestInbound(t *testing.T) {
	Convey("Given ", t, func() {
		dc := New()
		cs := make(chan struct{}, 100)
		Convey("When ", t, func() {
			// TODO need a name? can I use reflect?
			dc.Load(cs)

			cs := dc.MakeOutbound()
		})
	})
}

func TestOutbound(t *testing.T) {
	Convey("Given ", t, func() {
		d := New()
		cs := make(chan struct{}, 100)
		Convey("When ", t, func() {
			dc.Dump(cs)

			cs := dc.MakeInbound()
		})
	})
}

func TestObserver(t *testing.T) {
	Convey("Given ", t, func() {
		d := New()
		cs := make(chan struct{}, 100)
		Convey("When ", t, func() {
			dc.Load(cs)

			cs := dc.Make()
		})
	})
}
