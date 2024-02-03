package file

import (
	"context"
	"testing"
)

func TestFileWrite(t *testing.T) {
	fileName := "test.out"
	data := []byte("DOR")

	c, err := NewSink(fileName)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	out, err := c.Write(context.Background(), data)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	if string(data) != string(out) {
		t.Errorf("expected out to be %s got %s", data, out)
	}
}
