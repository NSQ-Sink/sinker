package nsq

import (
	"testing"

	"github.com/nsqio/go-nsq"
)

func TestModule_Run(t *testing.T) {
	type fields struct {
		nsqConsumer *nsq.Consumer
		source      []string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "test error empty consumer",
			fields: fields{
				nsqConsumer: nil,
			},
			wantErr: true,
		},
		{
			name: "test not error consumer while empty params",
			fields: fields{
				nsqConsumer: &nsq.Consumer{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Module{
				nsqConsumer: tt.fields.nsqConsumer,
			}
			if err := m.Run(); (err != nil) != tt.wantErr {
				t.Errorf("Module.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
