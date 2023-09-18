package streamer

// Run to run all handler in the consumer
func (m *NSQModule) Run() error {
	var err error

	// need to start all consumer
	for _, c := range m.consumers {
		if err = c.Run(); err != nil {
			return err
		}
	}

	return nil
}
