package event

import (
	"reflect"
	"testing"
)

func TestNewEvent(t *testing.T) {
	type args struct {
		topic  string
		source []string
	}
	tests := []struct {
		name string
		args args
		want Event
	}{
		{
			name: "test create event",
			args: args{
				topic:  "my_topic",
				source: []string{"dns1", "dns2"},
			},
			want: Module{
				topic:         "my_topic",
				sourceAddress: []string{"dns1", "dns2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEvent(tt.args.topic, tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}
