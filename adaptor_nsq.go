package dChannel

import (
	"time"

	nsq "github.com/nsqio/go-nsq"
	z "go.uber.org/zap"
)

type adaptor struct {
	delay  time.Duration
	logger *z.Logger
}

type adaptorNsq struct {
	adaptor

	config   *nsq.Config
	consumer *nsq.Consumer
	producer *nsq.Producer

	lookups []string
}

func NewNsq() *adaptorNsq {
	n := &adaptorNsq{}
	n.logger, _ = z.NewProduction()
	return n
}

func (n *adaptorNsq) Produce(name string, data []byte) {
	if err := n.producer.DeferredPublish(name, n.delay, data); err != nil {
		n.logger.Error("Connect problems occured when publishing.",
			z.String("name", name),
			z.ByteString("data", data))
		n.logger.Sync()
	}
}

func (n *adaptorNsq) Consume(name string, concurrency int, handler ConsumeHandler) {
	q, _ := nsq.NewConsumer(name, "ch", n.config)
	q.AddConcurrentHandlers(nsq.HandlerFunc(func(msg *nsq.Message) error {
		handler.HandleBytes(msg.Body)
		return nil
	}), concurrency)
	q.ConnectToNSQLookupds(n.lookups)
}
