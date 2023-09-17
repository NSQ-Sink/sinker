package file

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/nsqsink/sink/storage"
	"github.com/nsqsink/sink/storage/schema"
)

type File struct {
	file *os.File
}

func Open(fileName string) (storage.Storage, error) {
	f, err := os.Create(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open file because of %s", err.Error())
	}

	return &File{
		file: f,
	}, nil
}

func (f *File) Close(ctx context.Context) error {
	return f.file.Close()
}

func (f *File) Ping(ctx context.Context) error {
	_, err := f.file.Stat()
	return err
}

func (f *File) Write(ctx context.Context, data schema.MessageTransaction) error {
	d, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal data schema to json byte because of %s", err.Error())
	}

	buffer := new(bytes.Buffer)
	if err := json.Compact(buffer, d); err != nil {
		return fmt.Errorf("failed to json compact because of %s", err.Error())
	}

	if _, err = f.file.Write(buffer.Bytes()); err != nil {
		f.file.Close()
		return err
	}
	return nil
}

func (f *File) Read(ctx context.Context, filter schema.Filter) ([]schema.MessageTransaction, error) {
	scanner := bufio.NewScanner(f.file)

	var msgs []schema.MessageTransaction
	for scanner.Scan() {
		var msg schema.MessageTransaction
		if err := json.Unmarshal(scanner.Bytes(), &msg); err != nil {
			return nil, fmt.Errorf("failed to unmarshal log file because of %s", err.Error())
		}

		msgs = append(msgs, msg)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read log file because of %s", err.Error())
	}

	return msgs, nil
}
