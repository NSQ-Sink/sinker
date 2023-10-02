package streamer

import (
	"context"

	"github.com/nsqsink/sink/consumer"
	"github.com/nsqsink/sink/event"
	"github.com/nsqsink/sink/handler"
)

// RegisterConsumer implementation of register consumer method
// accepting event of the message, the handler for the event and the configuration of the consumer
func (m *NSQModule) RegisterConsumer(ctx context.Context, e event.Event, h handler.Handler, cfg consumer.Config) error {
	// create consumer
	c, err := consumer.New(ctx, e, h, cfg)
	if err != nil {
		return err
	}

	// adding consumer to list of consumer on streamer
	m.consumers = append(m.consumers, c)

	return nil
}
