package nsq

import (
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/nsqio/go-nsq"
	"github.com/nsqsink/sink/config"
	"github.com/nsqsink/sink/contract"
	message "github.com/nsqsink/sink/message/nsq"
)

type Module struct {
	nsqConsumer      *nsq.Consumer
	sourceNSQD       []string
	sourceNSQLookupd []string
}

// New return consumer module / object
// accepting event of the message, the handler for the event and the configuration of the consumer
func New(ctx context.Context, e contract.Event, h contract.Handler, cfg Config) (contract.Consumer, error) {
	// checking required data
	if e.GetTopic() == "" {
		return nil, errors.New("empty event topic")
	}

	if len(e.GetSource()) == 0 {
		return nil, errors.New("empty event source")
	}

	if h == nil {
		return nil, errors.New("empty handler")
	}

	// checking max attempt
	// using default value if empty
	if cfg.MaxAttempt <= 0 {
		cfg.MaxAttempt = constDefaultMaxAttempt
	}

	// checking max in flight
	// using default value if empty
	if cfg.MaxInFlight <= 0 {
		cfg.MaxInFlight = constDefaultMaxInflight
	}

	// generate random channel name from uuid if empty
	if cfg.ChannelName == "" {
		cfg.ChannelName = e.GetTopic() + "-" + uuid.NewString()
	}

	// create new consumer
	c, err := nsq.NewConsumer(e.GetTopic(), cfg.ChannelName, &nsq.Config{
		MaxAttempts: uint16(cfg.MaxAttempt),
		MaxInFlight: cfg.MaxInFlight,
	})
	if err != nil {
		return nil, err
	}

	// set log level
	c.SetLoggerLevel(toNSQLogLevel(cfg.LogLevel))

	// wrap handler to nsq handler
	handlerFn := func(msg *nsq.Message) error {
		return h.Handle(message.New(msg))
	}

	// add handler based on concurrent numbers
	if cfg.Concurrent > 0 {
		c.AddConcurrentHandlers(nsq.HandlerFunc(handlerFn), cfg.Concurrent)
	} else {
		c.AddHandler(nsq.HandlerFunc(handlerFn))
	}

	// parse source
	var (
		sourceNSQD       []string
		sourceNSQLookupd []string
	)

	for _, source := range e.GetSource() {
		if strings.Contains(source, config.ConstPrefixSourceNSQD) {
			sourceNSQD = append(sourceNSQD, strings.TrimLeft(source, config.ConstPrefixSourceNSQD))
			continue
		}

		if strings.Contains(source, config.ConstPrefixSourceNSQLookupd) {
			sourceNSQD = append(sourceNSQLookupd, strings.TrimLeft(source, config.ConstPrefixSourceNSQLookupd))
			continue
		}
	}

	// return consumer
	return Module{
		nsqConsumer:      c,
		sourceNSQD:       sourceNSQD,
		sourceNSQLookupd: sourceNSQLookupd,
	}, nil
}

// Run is a method to run / start the consumer to listen from an event
func (m Module) Run() error {
	if len(m.sourceNSQLookupd) == 0 && len(m.sourceNSQD) == 0 {
		return errors.New("empty source")
	}

	// run the consumer by connecting to nsqlookupd
	if len(m.sourceNSQLookupd) > 0 {
		if err := m.nsqConsumer.ConnectToNSQLookupds(m.sourceNSQLookupd); err != nil {
			return err
		}
	}

	if len(m.sourceNSQD) > 0 {
		if err := m.nsqConsumer.ConnectToNSQDs(m.sourceNSQD); err != nil {
			return err
		}
	}

	return nil
}

// Stop is a method to stop and close the consumer from listening an event
func (m Module) Stop() error {
	if m.nsqConsumer == nil {
		return errors.New("empty consumer")
	}

	// stop the consumer
	m.nsqConsumer.Stop()

	// wait until stopped (block)
	<-m.nsqConsumer.StopChan

	return nil
}
