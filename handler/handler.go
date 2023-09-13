package handler

import (
	"github.com/nsqsink/sink/message"
)

type Handler interface {
	Handle(msg message.Messager) error
}
