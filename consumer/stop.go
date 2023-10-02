package consumer

// Stop is a method to stop and close the consumer from listening an event
func (m Module) Stop() error {
	m.nsqConsumer.Stop()
	return nil
}
