package storage

import (
	"context"

	"github.com/nsqsink/sink/storage/schema"
)

type Storage interface {
	// Ping pings the database.
	Ping(ctx context.Context) error

	// Insert inserts the given data into the database.
	Write(ctx context.Context, data schema.MessageTransaction) error

	// Select selects the data from the database.
	Read(ctx context.Context, filter schema.Filter) ([]schema.MessageTransaction, error)

	// Close closes the database.
	Close(ctx context.Context) error
}
