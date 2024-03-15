package models

import "time"

type OPCDAHealthCheck struct {
	Server    string        `json:"server"`
	Nodes     []string      `json:"nodes"`
	Tags      []string      `json:"tags"`
	HeartBeat time.Duration `json:"heart_beat"`
}
