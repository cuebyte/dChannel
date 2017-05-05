package dChannel

const (
	Inbound = iota
	Outbound
)

type dChannel struct {
}

func (dc dChannel) MakeInt8(capacity uint, mode uint8) chan int8 {
	c := make(chan int8, capacity)
	return c
}
func (dc dChannel) MakeInt16(capacity uint, mode uint8) chan int16 {
	c := make(chan int16, capacity)
	return c
}
func (dc dChannel) MakeInt32(capacity uint, mode uint8) chan int32 {
	c := make(chan int32, capacity)
	return c
}
func (dc dChannel) MakeInt64(capacity uint, mode uint8) chan int64 {
	c := make(chan int64, capacity)
	return c
}
func (dc dChannel) MakeUint8(capacity uint, mode uint8) chan uint8 {
	c := make(chan uint8, capacity)
	return c
}
func (dc dChannel) MakeUint16(capacity uint, mode uint8) chan uint16 {
	c := make(chan uint16, capacity)
	return c
}
func (dc dChannel) MakeUint32(capacity uint, mode uint8) chan uint32 {
	c := make(chan uint32, capacity)
	return c
}
func (dc dChannel) MakeUint64(capacity uint, mode uint8) chan uint64 {
	c := make(chan uint64, capacity)
	return c
}
func (dc dChannel) MakeFloat32(capacity uint, mode uint8) chan float32 {
	c := make(chan float32, capacity)
	return c
}
func (dc dChannel) MakeFloat64(capacity uint, mode uint8) chan float64 {
	c := make(chan float64, capacity)
	return c
}
func (dc dChannel) MakeString(capacity uint, mode uint8) chan string {
	c := make(chan string, capacity)
	return c
}
func (dc dChannel) MakeBytes(capacity uint, mode uint8) chan []byte {
	c := make(chan []byte, capacity)
	return c
}
func (dc dChannel) MakeBool(capacity uint, mode uint8) chan bool {
	c := make(chan bool, capacity)
	return c
}
func (dc dChannel) MakeSignal(capacity uint, mode uint8) chan struct{} {
	c := make(chan struct{}, capacity)
	return c
}

func New() *dChannel {
	return &dChannel{}
}

func x() chan bool {
	return make(chan bool, 10)
}
