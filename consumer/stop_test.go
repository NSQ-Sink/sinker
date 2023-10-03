package consumer

import (
	"testing"

	"github.com/nsqio/go-nsq"
)

func TestModule_Stop(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Module{
				nsqConsumer: tt.fields.nsqConsumer,
				source:      tt.fields.source,
			}
			if err := m.Stop(); (err != nil) != tt.wantErr {
				t.Errorf("Module.Stop() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
