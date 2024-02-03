package nsq

import "github.com/nsqsink/sink/contract"

// NewEvent create new event
func New(topic string, sourceNSQD []string, sourceNSQLookupd []string) contract.Event {
	var source []string
	for _, temp := range sourceNSQD {
		if temp != "" {
			source = append(source, ConstPrefixSourceNSQD+temp)
		}
	}
	for _, temp := range sourceNSQLookupd {
		if temp != "" {
			source = append(source, ConstPrefixSourceNSQLookupd+temp)
		}
	}
	return Event{topic: topic, source: source}
}

// GetTopic return topic name
func (e Event) GetTopic() string {
	return e.topic
}

// GetSource return the source address for the topic
func (e Event) GetSource() []string {
	return e.source
}
