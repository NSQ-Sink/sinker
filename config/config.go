package config

import "strings"

// ParseSource parse source to array of string by comma
func (c *Consumer) ParseSource(sourceNSQD, sourceNSQLookupd string) {
	tempSourceNSQD := strings.Split(sourceNSQD, ",")
	for _, nsqdAddr := range tempSourceNSQD {
		if nsqdAddr != "" {
			c.Source.NSQD = append(c.Source.NSQD, nsqdAddr)
		}
	}

	tempSourceNSQLookupd := strings.Split(sourceNSQLookupd, ",")
	for _, nsqdAddr := range tempSourceNSQLookupd {
		if nsqdAddr != "" {
			c.Source.NSQLookupd = append(c.Source.NSQLookupd, nsqdAddr)
		}
	}
}
