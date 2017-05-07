package dChannel

import (
	"time"

	"go.uber.org/zap"
)

type Adaptor interface {
	Produce(name string, data []byte)
	Consume(name string, concurrency uint, handler ConsumeHandler) Consumer
}
type Consumer interface {
	Connect(addrs ...string)
	Disconnect(addrs ...string)
	Stop()
}
type ConsumeHandler interface {
	HandleBytes(data []byte) error
}
type ConsumeFunc func(data []byte) error            // impl ConsumeHandler
func (f ConsumeFunc) HandleBytes(data []byte) error { return f(data) }

type AbstractAdaptor struct {
	delay  time.Duration
	logger *zap.Logger
}
