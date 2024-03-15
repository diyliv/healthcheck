package models

import "context"

type Cache interface {
	Put(ctx context.Context, key interface{}, value interface{}) error
	Get(ctx context.Context, key interface{}) (interface{}, error)
	Remove(ctx context.Context, key interface{}) error
}
