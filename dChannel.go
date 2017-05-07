package dChannel

import (
	"fmt"

	"github.com/cuebyte/dChannel/pb"
	"github.com/gogo/protobuf/proto"
	"go.uber.org/zap"
)

type Connector interface {
	Produce(name string, data []byte)
	Consume(name string, concurrency uint, handler ConsumeHandler)
}

type ConsumeHandler interface {
	HandleBytes(data []byte) error
}
type ConsumeFunc func(data []byte) error // impl ConsumeHandler

type Consumer interface {
	Connect(addrs ...string)
	Disconnect(addrs ...string)
	Stop()
}

func (f ConsumeFunc) HandleBytes(data []byte) error { return f(data) }

type dChannel struct {
	conn Connector

	inCtrlMap   map[string]chan uint32
	outCtrlMap  map[string]chan uint32
	concurrency uint

	logger *zap.Logger
}

func New() *dChannel {
	dc := &dChannel{}
	return dc
}

func (dc *dChannel) PipeIn(name string, c interface{}) error {
	return dc.registerInbound(name, c)
}
func (dc *dChannel) PipeOut(name string, c interface{}) error {
	return dc.registerOutbound(name, c)
}
func (dc *dChannel) Close(name string) error {
	defer delete(dc.outCtrlMap, name)
	return nil
}

func (dc *dChannel) registerInbound(name string, c interface{}) error {
	if err := dc.regInCtrl(name); err != nil {
		return err
	}
	switch c.(type) {
	// case chan struct{}:
	case chan string:
		dc.regInString(name, c.(chan string))
	// case chan []byte:
	// case chan bool:
	// case chan int:
	// case chan uint:
	// case chan float32:
	// case chan float64:
	default:
	}
	return nil
}

// the fucking golang doesn't support overridding
func (dc *dChannel) registerOutbound(name string, c interface{}) error {
	if err := dc.regOutCtrl(name); err != nil {
		return err
	}

	switch c.(type) {
	// case chan struct{}:
	// dc.regStructOut(name, c.(chan struct{}))
	case chan string:
		dc.regOutString(name, c.(chan string))
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
	return nil
}

func (dc *dChannel) regInString(name string, c chan<- string) {
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

func (dc *dChannel) regOutString(name string, c <-chan string) {
	dc.arrange(func() {
		data, err := proto.Marshal(&pb.String{<-c})
		dc.marshalErrorHdlr(err, name, "String")
		dc.conn.Produce(name, data)
	})
}

func (dc *dChannel) regInCtrl(name string) error {
}

func (dc *dChannel) regOutCtrl(name string) error {
	if _, ok := dc.outCtrlMap[name]; !ok {
		return fmt.Errorf("register name duplicated. name=%s", name)
	}
	dc.outCtrlMap[name] = make(chan uint32)
	go func() {
		ctrlName := name + "-ctrl"
		data, err := proto.Marshal(&pb.Control{<-dc.outCtrlMap[name]})
		dc.marshalErrorHdlr(err, name, "Controll")
		dc.conn.Produce(ctrlName, data)
	}()
	return nil
}

func (dc *dChannel) marshalErrorHdlr(err error, name, typ string) {
	if err != nil {
		dc.logger.Error("error occured when outbound marshaling.",
			zap.String("name", name),
			zap.String("type", typ))
	}
}
