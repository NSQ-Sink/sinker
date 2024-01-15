package washtub

import (
	"context"

	"github.com/nsqsink/sink/entities"
)

type Washtuber interface {
	Pulse(ctx context.Context, data entities.PulseRequest) chan error

	Message(ctx context.Context, data entities.MessageRequest) (*entities.MessageResponse, error)
}
