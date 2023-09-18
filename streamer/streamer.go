package streamer

import (
	"github.com/nsqsink/sink/consumer"
)

// NSQModule struct
// struct for
type NSQModule struct {
	consumers []consumer.Consumer
}

// New
// return result initialization of NSQModule consumer
func New() Streamer {
	module := &NSQModule{
		consumers: make([]consumer.Consumer, 0),
	}

	return module
}
