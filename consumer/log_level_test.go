package consumer

import (
	"reflect"
	"testing"

	"github.com/nsqio/go-nsq"
)

func TestLogLevel_ToNSQLogLevel(t *testing.T) {
	tests := []struct {
		name string
		l    LogLevel
		want nsq.LogLevel
	}{
		{
			name: "test debug",
			l:    LogLevel("debug"),
			want: nsq.LogLevelDebug,
		},
		{
			name: "test info",
			l:    LogLevel("info"),
			want: nsq.LogLevelInfo,
		},
		{
			name: "test warn",
			l:    LogLevel("warn"),
			want: nsq.LogLevelWarning,
		},
		{
			name: "test error",
			l:    LogLevel("error"),
			want: nsq.LogLevelError,
		},
		{
			name: "test others",
			l:    LogLevel("others"),
			want: nsq.LogLevelMax,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.ToNSQLogLevel(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LogLevel.ToNSQLogLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
