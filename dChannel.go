package dChannel

const (
	inbound = iota
	outbound
)

type dChannel struct {
}

func (dc *dChannel) PipeIn(c interface{}) {
	dc.rigister(c, inbound)
}
func (dc *dChannel) PipeOut(c interface{}) {
	dc.rigister(c, outbound)
}
func (dc *dChannel) rigister(c interface{}, mode uint8) {
	switch c.(type) {
	case chan struct{}:
		c.(chan struct{})
	case chan string:
		c.(chan string)
	case chan []byte:
		c.(chan []byte)
	case chan bool:
		c.(chan bool)
	case chan int8:
		c.(chan int8)
	case chan int16:
		c.(chan int16)
	case chan int32:
		c.(chan int32)
	case chan int64:
		c.(chan int64)
	case chan uint8:
		c.(chan uint8)
	case chan uint16:
		c.(chan uint16)
	case chan uint32:
		c.(chan uint32)
	case chan uint64:
		c.(chan uint64)
	case chan float32:
		c.(chan float32)
	case chan float64:
		c.(chan float64)
	default:
	}
}

func New() *dChannel {
	dc := &dChannel{}
	return dc
}
