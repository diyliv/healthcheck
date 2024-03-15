package storage

import (
	"context"
	"sync"

	cacheerrors "github.com/diyliv/healthcheck/pkg/errors"
)

// Put
// Get
// Remove
// Cache eviction
type cache struct {
	mu      sync.RWMutex
	storage map[interface{}]interface{}
}

func NewCache() *cache {
	return &cache{storage: make(map[interface{}]interface{})}
}

func (c *cache) Put(ctx context.Context, key, value interface{}) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.storage[key]
	if ok {
		return cacheerrors.ErrAlreadyExistsInCache
	}
	c.storage[key] = value
	return nil
}

func (c *cache) Get(ctx context.Context, key interface{}) (interface{}, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.storage[key]
	if !ok {
		return nil, cacheerrors.ErrNoSuchedCachedValue
	}
	return val, nil
}

func (c *cache) Remove(ctx context.Context, key interface{}) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.storage[key]
	if !ok {
		return cacheerrors.ErrNoSuchedCachedValue
	}
	delete(c.storage, key)
	return nil
}
