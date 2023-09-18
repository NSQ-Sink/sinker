package event

// GetSource return the source address for the topic
func (e Module) GetSource() []string {
	return e.sourceAddress
}
