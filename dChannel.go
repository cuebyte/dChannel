package dChannel

import (
	"github.com/cuebyte/dChannel/pb"
	"github.com/gogo/protobuf/proto"
)

const (
	inbound = iota
	outbound
)

type Connector interface {
	Produce(name string, data []byte)
	Consume(name string, concurrency uint, handler ConsumeHandler)
}

type ConsumeHandler interface {
	HandleBytes(data []byte) error
}
type ConsumeFunc func(data []byte) error // impl ConsumeHandler

func (f ConsumeFunc) HandleBytes(data []byte) error { return f(data) }

type dChannel struct {
	conn Connector

	concurrency uint
}

func New() *dChannel {
	dc := &dChannel{}
	return dc
}

func (dc *dChannel) PipeIn(name string, c interface{}) error {
	dc.registerInbound(name, c)
	return nil
}
func (dc *dChannel) PipeOut(name string, c interface{}) error {
	dc.registerOutbound(name, c)
	return nil
}

func (dc *dChannel) registerInbound(name string, c interface{}) {
	switch c.(type) {
	// case chan struct{}:
	case chan string:
	// case chan []byte:
	// case chan bool:
	// case chan int:
	// case chan uint:
	// case chan float32:
	// case chan float64:
	default:
	}
}

// the fucking golang doesn't support overridding
func (dc *dChannel) registerOutbound(name string, c interface{}) {
	switch c.(type) {
	// case chan struct{}:
	// dc.regStructOut(name, c.(chan struct{}))
	case chan string:
		dc.regStringOut(name, c.(chan string))
	// case chan []byte:
	// 	dc.regBytesOut(name, c.(chan []byte))
	// case chan bool:
	// 	dc.regBoolOut(name, c.(chan bool))
	// case chan int:
	// 	dc.regIntOut(name, c.(chan int))
	// case chan uint:
	// 	dc.regUintOut(name, c.(chan uint))
	// case chan float32:
	// 	dc.regFloat32Out(name, c.(chan float32))
	// case chan float64:
	// 	dc.regFloat64Out(name, c.(chan float64))
	default:
	}
}

func (dc *dChannel) regStringIn(name string, c chan<- string) {
	dc.conn.Consume(name, dc.concurrency, ConsumeFunc(func(data []byte) error {
		tmp := &pb.String{}
		if err := proto.Unmarshal(data, tmp); err != nil {
			return err
		}
		c <- tmp.Value
		return nil
	}))
}

func (dc *dChannel) arrange(f func()) {
	for i := uint(0); i < dc.concurrency; i++ {
		go f()
	}
}

func (dc *dChannel) regStringOut(name string, c <-chan string) {
	dc.arrange(func() {
		data, _ := proto.Marshal(&pb.String{<-c})
		dc.conn.Produce(name, data)
	})
}
