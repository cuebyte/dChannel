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
			dc.PipeIn("x", cs) // chan<-
		})
	})
}

func TestOutbound(t *testing.T) {
	Convey("Given ", t, func() {
		dc := New()
		cs := make(chan struct{}, 100)
		Convey("When ", t, func() {
			dc.PipeOut("x", cs) // <-chan
		})
	})
}

func TestObserver(t *testing.T) {
	Convey("Given ", t, func() {
		d := New()
		cs := make(chan struct{}, 100)
		Convey("When ", t, func() {
			dc.Pipe(cs)
		})
	})
}
