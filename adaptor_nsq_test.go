package dChannel

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewNsq(t *testing.T) {
	Convey("Given ", t, func() {
		nsq := NewNsq()
	})
}

func TestProduce(t *testing.T) {
	Convey("Given ", t, func() {
		nsq := NewNsq()
		nsq.Produce("asdf", []byte("123"))
	})
}

func TestConsume(t *testing.T) {
	Convey("Given ", t, func() {
		nsq := NewNsq()
		nsq.Consume("asdf")
	})
}
