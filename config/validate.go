package config

import (
	"errors"
)

// Validate validate configuration, return false if config is not valid
func (c App) Validate() (bool, error) {
	// iterate checking on each configurations
	for _, cfg := range c.Consumers {
		// checking topic
		if cfg.Topic == "" {
			return false, errors.New("topic is empty")
		}

		// checking source
		if len(cfg.Source.NSQD) == 0 && len(cfg.Source.NSQLookupd) == 0 {
			return false, errors.New("source is empty")
		}

		// checking max attempt
		if cfg.MaxAttempt <= 0 {
			return false, errors.New("not valid max attempt")
		}

		// checking max in flight
		if cfg.MaxInFlight <= 0 {
			return false, errors.New("not valid max in flight")
		}

		// checking concurrent
		if cfg.Concurrent <= 0 {
			return false, errors.New("not valid concurrent")
		}

		// checking sinker type
		if cfg.Sinker.Type == "" {
			return false, errors.New("empty sinker type")
		}

		// checking parser type
		if cfg.Sinker.Parser.Type == "" {
			return false, errors.New("empty parser type")
		}

		// checking parser template
		if cfg.Sinker.Parser.Template == "" {
			return false, errors.New("empty parser template")
		}

		//@TODO: adding validation on parser template here
	}

	// if no cfg configurations then return not valid
	if len(c.Consumers) == 0 {
		return false, errors.New("no consumer configuration")
	}

	// return if everything is okay
	return true, nil
}
