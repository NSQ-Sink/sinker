package config

import (
	"errors"
	"fmt"
	"strings"
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
		if cfg.Source == "" && cfg.SourceNSQD == "" && cfg.SourceNSQLookupd == "" {
			return false, errors.New("source is empty")
		}

		listSource := cfg.ParseSource()
		for _, source := range listSource {
			if source == "" {
				return false, fmt.Errorf("not valid source %s", source)
			}
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

// ParseSource parse source to array of string by comma
func (c Consumer) ParseSource() []string {
	tempSource := strings.Split(c.Source, ",")

	tempSourceNSQD := strings.Split(c.SourceNSQD, ",")
	for _, nsqdAddr := range tempSourceNSQD {
		if nsqdAddr != "" {
			tempSource = append(tempSource, ConstPrefixSourceNSQD+nsqdAddr)
		}
	}

	tempSourceNSQLookupd := strings.Split(c.SourceNSQLookupd, ",")
	for _, nsqdAddr := range tempSourceNSQLookupd {
		if nsqdAddr != "" {
			tempSource = append(tempSource, ConstPrefixSourceNSQLookupd+nsqdAddr)
		}
	}

	return tempSource
}
