package models

import "time"

type OPCDAHealthCheck struct {
	Server    string
	Nodes     []string
	Tags      []string
	HeartBeat time.Duration
}
