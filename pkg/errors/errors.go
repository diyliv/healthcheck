package errors

import "errors"

var (
	ErrNoServerSpecified = errors.New("server must be specified for OPCDA health check")
	ErrNoNodeSpecified   = errors.New("node must be specified for OPCDA health check")
)
