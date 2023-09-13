package streamer

import (
	"context"

	"github.com/nsqio/go-nsq"
	"github.com/nsqsink/sink/handler"
	"github.com/nsqsink/sink/message"
)

// NSQModule struct
// struct for
type NSQModule struct {
	consumers []*nsq.Consumer
}

// New
// return result initialization of NSQModule consumer
func New() Streamer {
	module := &NSQModule{
		consumers: make([]*nsq.Consumer, 0),
	}

	return module
}

// RegisterConsumer implementation of register consumer method
// accepting event of the message, the handler for the event and the configuration of the consumer
func (m *NSQModule) RegisterConsumer(ctx context.Context, e Event, h handler.Handler, cfg ConsumerConfig) error {
	// create new consumer
	consumer, err := nsq.NewConsumer(e.GetTopic(), cfg.ChannelName, &nsq.Config{
		MaxAttempts: uint16(cfg.MaxAttempt),
		MaxInFlight: cfg.MaxInFlight,
	})
	if err != nil {
		return err
	}

	// set log level
	consumer.SetLoggerLevel(nsq.LogLevelError)

	handlerFn := func(msg *nsq.Message) error {
		return h.Handle(message.New(msg))
	}

	// add handler
	if cfg.Concurrent > 0 {
		consumer.AddConcurrentHandlers(nsq.HandlerFunc(handlerFn), cfg.Concurrent)
	} else {
		consumer.AddHandler(nsq.HandlerFunc(handlerFn))
	}

	return nil
}

// Run method
// method to run all handler in the consumer
func (m *NSQModule) Run() error {

	// need to start all consumer

	return nil
}

// Stop method
// method to stop all consumer handler in the consumer
func (m *NSQModule) Stop() error {

	// need to stop all handler

	// close consumer

	return nil
}
