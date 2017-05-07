package dChannel

import (
	"time"

	nsq "github.com/nsqio/go-nsq"
	"go.uber.org/zap"
)

type adaptor struct {
	delay  time.Duration
	logger *zap.Logger
}

type nsqAdaptor struct {
	adaptor

	Lookups  []string
	Config   *nsq.Config
	producer *nsq.Producer
}

type nsqConsumerAdaptor struct { // impl Consumer
	nsq.Consumer
}

func (c *nsqConsumerAdaptor) Connect(addrs ...string) {
	c.ConnectToNSQLookupds(addrs)
}
func (c *nsqConsumerAdaptor) Disconnect(addrs ...string) {
	for _, addr := range addrs {
		c.DisconnectFromNSQLookupd(addr)
	}
}
func (c *nsqConsumerAdaptor) Stop() {
	c.Stop()
}

func NewNsqAdaptorDefault() (n *nsqAdaptor, err error) {
	return NewNsqAdaptor(nsq.NewConfig(), 2*time.Second)
}

func NewNsqAdaptor(
	config *nsq.Config,
	delay time.Duration,
	lookups ...string) (n *nsqAdaptor, err error) {

	n = &nsqAdaptor{}
	n.Config = config
	n.producer, err = nsq.NewProducer("127.0.0.1:4145", config) // local nsqd
	if err != nil {
		return
	}
	n.delay = delay
	n.logger, err = zap.NewProduction()
	if err != nil {
		return
	}
	n.Lookups = append([]string{"127.0.0.1:4160"}, lookups...) // default with local nsqlookup
	return
}

func (n *nsqAdaptor) Produce(name string, data []byte) {
	if err := n.producer.DeferredPublish(name, n.delay, data); err != nil {
		n.logger.Error("Connect problems occured when publishing.",
			zap.String("name", name),
			zap.ByteString("data", data))
		n.logger.Sync()
	}
}

func (n *nsqAdaptor) Consume(name string, concurrency int, handler ConsumeHandler) {
	q, _ := nsq.NewConsumer(name, "ch", n.Config)
	q.AddConcurrentHandlers(nsq.HandlerFunc(func(msg *nsq.Message) error {
		return handler.HandleBytes(msg.Body)
	}), concurrency)
	q.ConnectToNSQLookupds(n.Lookups)
}
