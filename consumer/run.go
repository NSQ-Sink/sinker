package consumer

// Run is a method to run / start the consumer to listen from an event
func (m Module) Run() error {
	// run the consumer by connecting to nsqlookupd
	if err := m.nsqConsumer.ConnectToNSQDs(m.source); err != nil {
		return err
	}

	return nil
}
