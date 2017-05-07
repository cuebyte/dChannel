package dChannel

import (
	"time"

	nsq "github.com/nsqio/go-nsq"
	"go.uber.org/zap"
)

type nsqAdaptor struct {
	AbstractAdaptor

	Lookups  []string
	Config   *nsq.Config
	producer *nsq.Producer
}

type nsqConsumerAdaptor struct { // impl Consumer
	nc    *nsq.Consumer
	addrs []string
}

func (a *nsqConsumerAdaptor) Connect(addrs ...string) {
	a.nc.ConnectToNSQLookupds(addrs)
}
func (a *nsqConsumerAdaptor) Disconnect(addrs ...string) {
	for _, addr := range a.addrs {
		a.nc.DisconnectFromNSQLookupd(addr)
	}
}
func (a *nsqConsumerAdaptor) Stop() {
	a.nc.Stop()
}

func NewNsqAdaptorDefault() (n *nsqAdaptor, err error) {
	return NewNsqAdaptor(nsq.NewConfig(), 2*time.Second)
}

func NewNsqAdaptor(config *nsq.Config, delay time.Duration, lookups ...string) (n *nsqAdaptor, err error) {
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
	if err := n.producer.DeferredPublish("dC:"+name, n.delay, data); err != nil {
		n.logger.Error("Connect problems occured when publishing.",
			zap.String("name", name),
			zap.ByteString("data", data))
		n.logger.Sync()
	}
}

func (n *nsqAdaptor) Consume(name string, concurrency int, handler ConsumeHandler) Consumer {
	q, _ := nsq.NewConsumer("dC:"+name, "ch", n.Config)
	q.AddConcurrentHandlers(nsq.HandlerFunc(func(msg *nsq.Message) error {
		return handler.HandleBytes(msg.Body)
	}), concurrency)
	q.ConnectToNSQLookupds(n.Lookups)
	return &nsqConsumerAdaptor{q, n.Lookups}
}
