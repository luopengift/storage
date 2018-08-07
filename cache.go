package storage

import (
	"context"
	"errors"
	"sync"
)

// Cache implement storager interface
type Cache struct {
	sync.Map
}

// CacheInit init cache
func CacheInit() (Storager, error) {
	return &Cache{
		Map: sync.Map{},
	}, nil
}

// Interface interface
func (c *Cache) Interface() interface{} {
	return &(c.Map)
}

// Put put
func (c *Cache) Put(ctx context.Context, key, val string, opts ...interface{}) (interface{}, error) {
	c.Map.Store(key, val)
	return nil, nil
}

// Get get
func (c *Cache) Get(ctx context.Context, key string, opts ...interface{}) (interface{}, error) {
	if v, ok := c.Map.Load(key); ok {
		return v, nil
	}
	return nil, errors.New("invalid value")
}

// Delete delete
func (c *Cache) Delete(ctx context.Context, key string, opts ...interface{}) (interface{}, error) {
	c.Map.Delete(key)
	return nil, nil
}

// Do do
func (c *Cache) Do(ctx context.Context, op interface{}) (interface{}, error) {
	return nil, nil
}

// Close close
func (c *Cache) Close() error {
	return nil
}
