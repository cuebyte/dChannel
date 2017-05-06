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

type protoType interface {
	Reset()
	String() string
	ProtoMessage()
	Descriptor() ([]byte, []int)
}

type dChannel struct {
	conn *Connector
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
func (dc *dChannel) Conn() *Connector { return dc.conn }

// the fucking golang doesn't support overridding
func (dc *dChannel) register(name string, c interface{}, mode uint8) {
	if mode == inbound {
		switch c.(type) {
		case chan struct{}:
			dc.regStructIn(name, c.(chan struct{}))
		case chan string:
			dc.regStringIn(name, c.(chan string))
		case chan []byte:
			dc.regBytesIn(name, c.(chan []byte))
		case chan bool:
			dc.regBoolIn(name, c.(chan bool))
		case chan int8:
			dc.regI8In(name, c.(chan int8))
		case chan int16:
			dc.regI16In(name, c.(chan int16))
		case chan int32:
			dc.regI32In(name, c.(chan int32))
		case chan int64:
			dc.regI64In(name, c.(chan int64))
		case chan uint8:
			dc.regU8In(name, c.(chan uint8))
		case chan uint16:
			dc.regU16In(name, c.(chan uint16))
		case chan uint32:
			dc.regU32In(name, c.(chan uint32))
		case chan uint64:
			dc.regU64In(name, c.(chan uint64))
		case chan float32:
			dc.regF32In(name, c.(chan float32))
		case chan float64:
			dc.regF64In(name, c.(chan float64))
		default:
		}
		if mode == outbound {
			switch c.(type) {
			case chan struct{}:
				go func(name string, ch <-chan struct{}) {
					data, _ := proto.Marshal(&Struct{true})
					dc.conn.Produce(name, data)
				}(name, c.(chan struct{}))
			case chan string:
				go func(name string, ch <-chan string) {
					data, _ := proto.Marshal(&String{<-ch})
					dc.conn.Produce(name, data)
				}(name, c.(chan string))
			case chan []byte:
				go func(name string, ch <-chan []byte) {
					data, _ := proto.Marshal(&Bytes{<-ch})
					dc.conn.Produce(name, data)
				}(name, c.(chan []byte))
			case chan bool:
				go func(name string, ch <-chan bool) {
					data, _ := proto.Marshal(&Bool{<-ch})
					dc.conn.Produce(name, data)
				}(name, c.(chan bool))
			case chan int8:
				go func(name string, ch <-chan int8) {
					data, _ := proto.Marshal(&Int8{int32(<-ch)})
					dc.conn.Produce(name, data)
				}(name, c.(chan int8))
			case chan int16:
				go func(name string, ch <-chan int16) {
					data, _ := proto.Marshal(&Int16{int32(<-ch)})
					dc.conn.Produce(name, data)
				}(name, c.(chan int16))
			case chan int32:
				go func(name string, ch <-chan int32) {
					data, _ := proto.Marshal(&Int32{<-ch})
					dc.conn.Produce(name, data)
				}(name, c.(chan int32))
			case chan int64:
				go func(name string, ch <-chan int64) {
					data, _ := proto.Marshal(&Int64{<-ch})
					dc.conn.Produce(name, data)
				}(name, c.(chan int64))
			case chan uint8:
				go func(name string, ch <-chan uint8) {
					data, _ := proto.Marshal(&Uint8{uint32(<-ch)})
					dc.conn.Produce(name, data)
				}(name, c.(chan uint8))
			case chan uint16:
				go func(name string, ch <-chan uint16) {
					data, _ := proto.Marshal(&Uint16{uint32(<-ch)})
					dc.conn.Produce(name, data)
				}(name, c.(chan uint16))
			case chan uint32:
				go func(name string, ch <-chan uint32) {
					data, _ := proto.Marshal(&Uint32{<-ch})
					dc.conn.Produce(name, data)
				}(name, c.(chan uint32))
			case chan uint64:
				go func(name string, ch <-chan uint64) {
					data, _ := proto.Marshal(&Uint64{<-ch})
					dc.conn.Produce(name, data)
				}(name, c.(chan uint64))
			case chan float32:
				go func(name string, ch <-chan float32) {
					data, _ := proto.Marshal(&Float32{<-ch})
					dc.conn.Produce(name, data)
				}(name, c.(chan float32))
			case chan float64:
				go func(name string, ch <-chan float64) {
					data, _ := proto.Marshal(&Float64{<-ch})
					dc.conn.Produce(name, data)
				}(name, c.(chan float64))
			default:
			}
		}
	}

}
func (dc *dChannel) regStructIn(name string, c chan<- struct{}) {}
func (dc *dChannel) regStringIn(name string, c chan<- string) {

}
func (dc *dChannel) regBytesIn(name string, c chan<- []byte) {}
func (dc *dChannel) regBoolIn(name string, c chan<- bool)    {}
func (dc *dChannel) regI8In(name string, c chan<- int8)      {}
func (dc *dChannel) regI16In(name string, c chan<- int16)    {}
func (dc *dChannel) regI32In(name string, c chan<- int32)    {}
func (dc *dChannel) regI64In(name string, c chan<- int64)    {}
func (dc *dChannel) regU8In(name string, c chan<- uint8)     {}
func (dc *dChannel) regU16In(name string, c chan<- uint16)   {}
func (dc *dChannel) regU32In(name string, c chan<- uint32)   {}
func (dc *dChannel) regU64In(name string, c chan<- uint64)   {}
func (dc *dChannel) regF32In(name string, c chan<- float32)  {}
func (dc *dChannel) regF64In(name string, c chan<- float64)  {}

func (dc *dChannel) regStructOut(name string, c <-chan struct{}) {}
func (dc *dChannel) regStringOut(name string, c <-chan string) {
	go func() {
		data, _ := proto.Marshal(&String{<-c})
		dc.conn.Produce(name, data)
	}()
}
func (dc *dChannel) regBytesOut(name string, c <-chan []byte) {}
func (dc *dChannel) regBoolOut(name string, c <-chan bool)    {}
func (dc *dChannel) regI8Out(name string, c <-chan int8)      {}
func (dc *dChannel) regI16Out(name string, c <-chan int16)    {}
func (dc *dChannel) regI32Out(name string, c <-chan int32)    {}
func (dc *dChannel) regI64Out(name string, c <-chan int64)    {}
func (dc *dChannel) regU8Out(name string, c <-chan uint8)     {}
func (dc *dChannel) regU16Out(name string, c <-chan uint16)   {}
func (dc *dChannel) regU32Out(name string, c <-chan uint32)   {}
func (dc *dChannel) regU64Out(name string, c <-chan uint64)   {}
func (dc *dChannel) regF32Out(name string, c <-chan float32)  {}
func (dc *dChannel) regF64Out(name string, c <-chan float64)  {}
