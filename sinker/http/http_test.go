package http

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHTTPWrite(t *testing.T) {
	expected := "PONG!"

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, expected)
	}))
	defer svr.Close()

	c, err := NewSink(svr.URL, http.MethodPost,
		WithTimeout(1*time.Second),
		WithHeader(map[string]interface{}{
			"key": "value",
		}),
		WithRetry(RetryConfig{
			MaxRetries:      3,
			Backoff:         2 * time.Second,
			RetryStatusCode: []int{500, 502},
		}),
	)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	res, err := c.Write(context.Background(), []byte("PING"))
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	if string(res) != expected {
		t.Errorf("expected res to be %s got %s", expected, res)
	}
}
