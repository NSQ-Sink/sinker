package consumer

// Config config for consumer
type Config struct {
	ChannelName string // name of the consumer channel
	Concurrent  int    // number of concurrent consumer
	MaxAttempt  int    // max attempt of consumer to handle a message
	MaxInFlight int
	LogLevel    LogLevel // setting for log level (1 - )
}

const (
	constDefaultMaxAttempt  int = 3
	constDefaultMaxInflight int = 10
)
