package nsq

import (
	"reflect"
	"testing"

	"github.com/nsqio/go-nsq"
	"github.com/nsqsink/sink/log"
)

func TestLogLevel_toNSQLogLevel(t *testing.T) {
	tests := []struct {
		name string
		l    log.LogLevel
		want nsq.LogLevel
	}{
		{
			name: "test debug",
			l:    log.LogLevel("debug"),
			want: nsq.LogLevelDebug,
		},
		{
			name: "test info",
			l:    log.LogLevel("info"),
			want: nsq.LogLevelInfo,
		},
		{
			name: "test warn",
			l:    log.LogLevel("warn"),
			want: nsq.LogLevelWarning,
		},
		{
			name: "test error",
			l:    log.LogLevel("error"),
			want: nsq.LogLevelError,
		},
		{
			name: "test others",
			l:    log.LogLevel("others"),
			want: nsq.LogLevelMax,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toNSQLogLevel(tt.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LogLevel.ToNSQLogLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
