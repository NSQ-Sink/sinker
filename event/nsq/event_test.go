package nsq

import (
	"reflect"
	"testing"
)

func TestNewEvent(t *testing.T) {
	type args struct {
		topic            string
		sourceNSQD       []string
		sourceNSQLookupd []string
	}
	tests := []struct {
		name string
		args args
		want Event
	}{
		{
			name: "test create event",
			args: args{
				topic:            "my_topic",
				sourceNSQD:       []string{"dns1", "dns2"},
				sourceNSQLookupd: []string{"dns3", "dns4"},
			},
			want: Event{
				topic:  "my_topic",
				source: []string{"nsqd_dns1", "nsqd_dns2", "nsqlookupd_dns3", "nsqlookupd_dns4"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.topic, tt.args.sourceNSQD, tt.args.sourceNSQLookupd); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}
