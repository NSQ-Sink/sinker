package streamer

import (
	"sync"

	"github.com/nsqsink/sink/consumer"
)

// Stop to stop all consumer handler in the consumer
// using go routine to make it faster
func (m *NSQModule) Stop() error {
	var (
		err error
		wg  sync.WaitGroup
		mux sync.Mutex
	)

	// add wait group based on list of consumer registered
	wg.Add(len(m.consumers))

	for _, c := range m.consumers {
		go func(nsqConsumer *consumer.Consumer) {
			defer wg.Done()

		}(c)
	}

	return nil
}
