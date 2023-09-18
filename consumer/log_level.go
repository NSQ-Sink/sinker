package consumer

import "github.com/nsqio/go-nsq"

// LogLevel specifies the severity of a given log message
type LogLevel string

// Log levels
const (
	LogLevelDebug   LogLevel = "debug"
	LogLevelInfo    LogLevel = "info"
	LogLevelWarning LogLevel = "warn"
	LogLevelError   LogLevel = "error"
)

// ToNSQLogLevel return log level for NSQ
func (l LogLevel) ToNSQLogLevel() nsq.LogLevel {
	switch l {
	case "debug":
		return nsq.LogLevelDebug
	case "info":
		return nsq.LogLevelInfo
	case "warn":
		return nsq.LogLevelWarning
	case "error":
		return nsq.LogLevelError
	}

	return nsq.LogLevelMax
}
