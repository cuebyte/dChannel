package dChannel

type dChannel struct {
}

func (dc *dChannel) PipeInSignal(chan<- struct{}) {}
func (dc *dChannel) PipeInString(chan<- string)   {}
func (dc *dChannel) PipeInBytes(chan<- []byte)    {}
func (dc *dChannel) PipeInBool(chan<- bool)       {}
func (dc *dChannel) PipeInInt8(chan<- int8)       {}
func (dc *dChannel) PipeInInt16(chan<- int16)     {}
func (dc *dChannel) PipeInInt32(chan<- int32)     {}
func (dc *dChannel) PipeInInt64(chan<- int64)     {}
func (dc *dChannel) PipeInUint8(chan<- uint8)     {}
func (dc *dChannel) PipeInUint16(chan<- uint16)   {}
func (dc *dChannel) PipeInUint32(chan<- uint32)   {}
func (dc *dChannel) PipeInUint64(chan<- uint64)   {}
func (dc *dChannel) PipeInFloat32(chan<- float32) {}
func (dc *dChannel) PipeInFloat64(chan<- float64) {}

func (dc *dChannel) PipeOutSignal(<-chan struct{}) {}
func (dc *dChannel) PipeOutString(<-chan string)   {}
func (dc *dChannel) PipeOutBytes(<-chan []byte)    {}
func (dc *dChannel) PipeOutBool(<-chan bool)       {}
func (dc *dChannel) PipeOutInt8(<-chan int8)       {}
func (dc *dChannel) PipeOutInt16(<-chan int16)     {}
func (dc *dChannel) PipeOutInt32(<-chan int32)     {}
func (dc *dChannel) PipeOutInt64(<-chan int64)     {}
func (dc *dChannel) PipeOutUint8(<-chan uint8)     {}
func (dc *dChannel) PipeOutUint16(<-chan uint16)   {}
func (dc *dChannel) PipeOutUint32(<-chan uint32)   {}
func (dc *dChannel) PipeOutUint64(<-chan uint64)   {}
func (dc *dChannel) PipeOutFloat32(<-chan float32) {}
func (dc *dChannel) PipeOutFloat64(<-chan float64) {}

func New() *dChannel {
	dc := &dChannel{}
	return dc
}
