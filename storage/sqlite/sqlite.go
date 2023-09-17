package sqlite

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/nsqsink/sink/storage"
	"github.com/nsqsink/sink/storage/schema"
)

type Sqlite struct {
	db *sqlx.DB
}

func Open(fileName string) (storage.Storage, error) {
	// Connect a new database connection.
	db, err := sqlx.Open("sqlite3", fileName)
	if err != nil {
		return nil, err
	}

	return &Sqlite{
		db: db,
	}, nil
}

func (s *Sqlite) Close(ctx context.Context) error {
	return s.db.Close()
}

func (s *Sqlite) Ping(ctx context.Context) error {
	return s.db.Ping()
}

func (s *Sqlite) Write(ctx context.Context, data schema.MessageTransaction) error {
	return nil
}

func (s *Sqlite) Read(ctx context.Context, filter schema.Filter) ([]schema.MessageTransaction, error) {
	return nil, nil
}
