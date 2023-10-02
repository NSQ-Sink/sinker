package event

type Module struct {
	topic         string   // topic name
	sourceAddress []string // source of the topic, for nsq its a nsqlookupd address
}

// NewEvent create new event
func NewEvent(topic string, source []string) Event {
	return Module{topic: topic, sourceAddress: source}
}
