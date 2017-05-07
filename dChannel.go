package dChannel

import (
	"fmt"

	"github.com/cuebyte/dChannel/pb"
	"github.com/gogo/protobuf/proto"
	"go.uber.org/zap"
)

type dChannel struct {
	conn Adaptor

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
	if _, ok := dc.inCtrlMap[name]; !ok {
		return fmt.Errorf("register name duplicated. name=%s", name)
	}
	var consumer Consumer
	switch c.(type) {
	// case chan struct{}:
	case chan string:
		consumer = dc.regInString(name, c.(chan string))
	// case chan []byte:
	// case chan bool:
	// case chan int:
	// case chan uint:
	// case chan float32:
	// case chan float64:
	default:
	}
	dc.regInCtrl(name, consumer)
	return nil
}

func (dc *dChannel) registerOutbound(name string, c interface{}) error {
	if _, ok := dc.outCtrlMap[name]; !ok {
		return fmt.Errorf("register name duplicated. name=%s", name)
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
	dc.regOutCtrl(name)
	return nil
}

func (dc *dChannel) regInString(name string, c chan<- string) Consumer {
	return dc.conn.Consume(name, dc.concurrency, ConsumeFunc(func(data []byte) error {
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

func (dc *dChannel) regInCtrl(name string, consumer Consumer) {
	dc.inCtrlMap[name] = make(chan uint32)
	ctrlName := name + "-ctrl"
	ctrlConsumer := dc.conn.Consume(ctrlName, 1, ConsumeFunc(func(data []byte) error {
		tmp := &pb.Control{}
		if err := proto.Unmarshal(data, tmp); err != nil {
			return err
		}
		dc.inCtrlMap[name] <- tmp.Value
		return nil
	}))
	go func() {
		// TODO concurrency control
		<-dc.inCtrlMap[name]
		// TODO analysis the output
		consumer.Stop()
		ctrlConsumer.Stop()
		delete(dc.inCtrlMap, name)
	}()
}

func (dc *dChannel) regOutCtrl(name string) {
	dc.outCtrlMap[name] = make(chan uint32)
	ctrlName := name + "-ctrl"
	go func() {
		data, err := proto.Marshal(&pb.Control{<-dc.outCtrlMap[name]})
		dc.marshalErrorHdlr(err, name, "Controll")
		dc.conn.Produce(ctrlName, data)
	}()
}

func (dc *dChannel) marshalErrorHdlr(err error, name, typ string) {
	if err != nil {
		dc.logger.Error("error occured when outbound marshaling.",
			zap.String("name", name),
			zap.String("type", typ))
	}
}
