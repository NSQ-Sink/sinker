package streamer

import "github.com/nsqio/go-nsq"

type Consumer interface {
	Run()
}

// ConsumerConfig config for consumer
type ConsumerConfig struct {
	ChannelName string // name of the consumer channel
	Concurrent  int    // number of concurrent consumer
	MaxAttempt  int    // max attempt of consumer to handle a message
	MaxInFlight int
}

type ConsumerModule struct {
	c      nsq.Consumer
	source string
}

// func NewConsumer() Consumer {

// }
