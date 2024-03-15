package errors

import "errors"

var (
	ErrAlreadyExistsInCache = errors.New("such value already exists")
	ErrNoSuchedCachedValue  = errors.New("no such value in cache")
	ErrOPCDAConvertion      = errors.New("could not convert interface{} to opc.Connection")
	ErrNoServerSpecified    = errors.New("server must be specified for OPCDA health check")
	ErrNoNodeSpecified      = errors.New("node must be specified for OPCDA health check")
)
