package washtubhttp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/nsqsink/sink/entity"
	"github.com/nsqsink/sink/washtub"
)

const (
	DefaultTimeout       = 2 * time.Second
	DefaultPulseInterval = 5 * time.Second
)

type Client struct {
	// The `stopPulse` channel of type `chan struct{}` is used to signal the client to stop pulsing. It is
	// a synchronization mechanism that allows the client to gracefully stop its pulsing operation. When a
	// value is sent on this channel, it indicates that the client should stop pulsing and exit.
	stopPulse chan struct{}

	// The `pulseInterval time.Duration` is a field in the `Client` struct that represents the interval at
	// which the client should send pulses. It is of type `time.Duration`, which allows you to specify the
	// duration of the interval in a human-readable format, such as seconds, minutes, or hours. This field
	// determines how often the client should send pulses to a server or perform some other periodic task.
	pulseInterval time.Duration

	// The `address string` field in the `Client` struct is used to store the address of the server that
	// the client will send requests to. It represents the URL or IP address of the server. This field is
	// used in the `call` method to create an HTTP request to the specified server address.
	address string

	// client is the HTTP client to use.
	client *http.Client
}

func NewWashtuber(ctx context.Context, address string, options ...Option) (washtub.Washtuber, error) {
	c := &Client{
		client: &http.Client{
			Timeout: DefaultTimeout,
		},
		stopPulse:     make(chan struct{}, 1),
		pulseInterval: DefaultPulseInterval,
		address:       address,
	}

	for _, option := range options {
		if err := option(c); err != nil {
			return c, fmt.Errorf("failed to register option to Washtub Client because %s", err.Error())
		}
	}

	return c, nil
}

func (c *Client) Pulse(ctx context.Context, data entity.PulseRequest) chan error {
	var (
		errCh  = make(chan error, 2)
		ticker = time.NewTicker(c.pulseInterval)
	)
	err := c.call(ctx, data)
	if err != nil {
		errCh <- err
	}
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
			case <-c.stopPulse:
				return
			}
			err := c.call(ctx, data)
			if err != nil {
				select {
				case errCh <- err:
				default:
				}
			}
		}
	}()

	return errCh
}

func (c *Client) Message(ctx context.Context, data entity.MessageRequest) (*entity.MessageResponse, error) {
	return &entity.MessageResponse{}, nil
}

func (c *Client) call(ctx context.Context, body interface{}) error {
	b, err := json.Marshal(body)
	if err != nil {
		return err
	}

	bodyReader := bytes.NewReader(b)
	req, err := http.NewRequest(http.MethodPost, c.address+"/worker/pulse", bodyReader)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req = req.WithContext(ctx)
	req.Close = true
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var pr entity.PulseResponse
	if err := json.Unmarshal(response, &pr); err != nil {
		return err
	}

	return nil
}
