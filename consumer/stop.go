package consumer

import "errors"

// Stop is a method to stop and close the consumer from listening an event
func (m Module) Stop() error {
	if m.nsqConsumer == nil {
		return errors.New("consumer is nil")
	}

	m.nsqConsumer.Stop()

	return nil
}
