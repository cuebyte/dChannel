package dChannel

const (
	inbound = iota
	outbound
)

type dChannel struct {
}

func New() *dChannel {
	dc := &dChannel{}
	return dc
}

func (dc *dChannel) PipeIn(c interface{}) {
	dc.reg(c, inbound)
}
func (dc *dChannel) PipeOut(c interface{}) {
	dc.reg(c, outbound)
}
func (dc *dChannel) reg(c interface{}, mode uint8) {
	switch c.(type) {
	case chan struct{}:
		dc.regStruct(c.(chan struct{}), mode)
	case chan string:
		dc.regString(c.(chan string), mode)
	case chan []byte:
		dc.regBytes(c.(chan []byte), mode)
	case chan bool:
		dc.regBool(c.(chan bool), mode)
	case chan int8:
		dc.regI8(c.(chan int8), mode)
	case chan int16:
		dc.regI16(c.(chan int16), mode)
	case chan int32:
		dc.regI32(c.(chan int32), mode)
	case chan int64:
		dc.regI64(c.(chan int64), mode)
	case chan uint8:
		dc.regU8(c.(chan uint8), mode)
	case chan uint16:
		dc.regU16(c.(chan uint16), mode)
	case chan uint32:
		dc.regU32(c.(chan uint32), mode)
	case chan uint64:
		dc.regU64(c.(chan uint64), mode)
	case chan float32:
		dc.regF32(c.(chan float32), mode)
	case chan float64:
		dc.regF64(c.(chan float64), mode)
	default:
	}
}
func (dc *dChannel) regStruct(c chan struct{}, mode uint8) {}
func (dc *dChannel) regString(c chan string, mode uint8)   {}
func (dc *dChannel) regBytes(c chan []byte, mode uint8)    {}
func (dc *dChannel) regBool(c chan bool, mode uint8)       {}
func (dc *dChannel) regI8(c chan int8, mode uint8)         {}
func (dc *dChannel) regI16(c chan int16, mode uint8)       {}
func (dc *dChannel) regI32(c chan int32, mode uint8)       {}
func (dc *dChannel) regI64(c chan int64, mode uint8)       {}
func (dc *dChannel) regU8(c chan uint8, mode uint8)        {}
func (dc *dChannel) regU16(c chan uint16, mode uint8)      {}
func (dc *dChannel) regU32(c chan uint32, mode uint8)      {}
func (dc *dChannel) regU64(c chan uint64, mode uint8)      {}
func (dc *dChannel) regF32(c chan float32, mode uint8)     {}
func (dc *dChannel) regF64(c chan float64, mode uint8)     {}
