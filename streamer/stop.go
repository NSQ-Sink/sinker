package streamer

import (
	"github.com/nsqsink/sink/consumer"
)

// Stop to stop all consumer handler in the consumer
// using go routine to make it faster
func (m *NSQModule) Stop() error {
	errChan := make(chan error)

	for _, c := range m.consumers {
		go func(nsqConsumer *consumer.Consumer) {
			var err error
			defer func() {
				errChan <- err
			}()

			err = c.Stop()
		}(&c)
	}

	close(errChan)

	var err error
	for tempErr := range errChan {
		if tempErr != nil {
			err = tempErr
		}
	}

	return err
}
