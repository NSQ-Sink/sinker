package washtubhttp

import "time"

type Option func(*Client) error

func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) error {
		c.client.Timeout = timeout
		return nil
	}
}

func WithPulseInterval(interval time.Duration) Option {
	return func(c *Client) error {
		c.pulseInterval = interval
		return nil
	}
}
