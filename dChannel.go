package dChan

import "github.com/gogo/protobuf/proto"

const (
	inbound = iota
	outbound
)

type Connector interface {
	Produce(name string, data []byte)
	Consume(name string, data []byte)
}

type dChannel struct {
	conn Connector

	routineNumber int
}

func New() *dChannel {
	dc := &dChannel{}
	return dc
}

func (dc *dChannel) PipeIn(name string, c interface{}) {
	dc.register(name, c, inbound)
}
func (dc *dChannel) PipeOut(name string, c interface{}) {
	dc.register(name, c, outbound)
}

func (dc *dChannel) costomRegister(name string) {

}

// the fucking golang doesn't support overridding
func (dc *dChannel) register(name string, c interface{}, mode uint8) {
	if mode == inbound {
		switch c.(type) {
		case chan struct{}:
		case chan string:
		case chan []byte:
		case chan bool:
		case chan int:
		case chan uint:
		case chan float32:
		case chan float64:
		default:
		}
	}
	if mode == outbound {
		switch c.(type) {
		case chan struct{}:
			dc.regStructOut(name, c.(chan struct{}))
		case chan string:
			dc.regStringOut(name, c.(chan string))
		case chan []byte:
			dc.regBytesOut(name, c.(chan []byte))
		case chan bool:
			dc.regBoolOut(name, c.(chan bool))
		case chan int:
			dc.regIntOut(name, c.(chan int))
		case chan uint:
			dc.regUintOut(name, c.(chan uint))
		case chan float32:
			dc.regFloat32Out(name, c.(chan float32))
		case chan float64:
			dc.regFloat64Out(name, c.(chan float64))
		default:
		}
	}
}

func (dc *dChannel) arrange(f func()) {
	for i := 0; i < dc.routineNumber; i++ {
		go f()
	}
}

func (dc *dChannel) regStructIn(name string, c chan<- struct{}) {
	dc.arrange()
}
func (dc *dChannel) regStringIn(name string, c chan<- string) {
	dc.arrange()
}
func (dc *dChannel) regBytesIn(name string, c chan<- []byte) {
	dc.arrange()
}
func (dc *dChannel) regBoolIn(name string, c chan<- bool) {
	dc.arrange()
}
func (dc *dChannel) regIntIn(name string, c chan<- int8) {
	dc.arrange()
}
func (dc *dChannel) regUintIn(name string, c chan<- int8) {
	dc.arrange()
}
func (dc *dChannel) regFloat32In(name string, c chan<- float32) {
	dc.arrange()
}
func (dc *dChannel) regFloat64In(name string, c chan<- float64) {
	dc.arrange()
}

func (dc *dChannel) regStructOut(name string, c <-chan struct{}) {
	dc.arrange(func() {
		data, _ := proto.Marshal(&Struct{true})
		dc.conn.Produce(name, data)
	})
}
func (dc *dChannel) regStringOut(name string, c <-chan string) {
	dc.arrange(func() {
		data, _ := proto.Marshal(&String{<-c})
		dc.conn.Produce(name, data)
	})
}
func (dc *dChannel) regBytesOut(name string, c <-chan []byte) {
	dc.arrange(func() {
		data, _ := proto.Marshal(&Bytes{<-c})
		dc.conn.Produce(name, data)
	})
}
func (dc *dChannel) regBoolOut(name string, c <-chan bool) {
	dc.arrange(func() {
		data, _ := proto.Marshal(&Bool{<-c})
		dc.conn.Produce(name, data)
	})
}
func (dc *dChannel) regIntOut(name string, c <-chan int) {
	dc.arrange(func() {
		data, _ := proto.Marshal(&Int{int32(<-c)})
		dc.conn.Produce(name, data)
	})
}
func (dc *dChannel) regUintOut(name string, c <-chan uint) {
	dc.arrange(func() {
		data, _ := proto.Marshal(&Uint{uint32(<-c)})
		dc.conn.Produce(name, data)
	})
}
func (dc *dChannel) regFloat32Out(name string, c <-chan float32) {
	dc.arrange(func() {
		data, _ := proto.Marshal(&Float32{<-c})
		dc.conn.Produce(name, data)
	})
}
func (dc *dChannel) regFloat64Out(name string, c <-chan float64) {
	dc.arrange(func() {
		data, _ := proto.Marshal(&Float64{<-c})
		dc.conn.Produce(name, data)
	})
}
