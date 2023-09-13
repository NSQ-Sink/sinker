package streamer

import (
	"context"

	"github.com/nsqsink/sink/handler"
)

// Streamer is an interface which can be implemented
// to run all consumer that already registered to the streamer
type Streamer interface {
	// RegisterConsumer method
	// method to register consumer to the streamer
	RegisterConsumer(ctx context.Context, e Event, h handler.Handler, cfg ConsumerConfig) error

	// Run method
	// method to run all consumer in the streamer
	Run() error

	// Stop method
	// method to stop all consumer in the streamer
	Stop() error
}
