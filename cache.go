package storage

import (
	"context"
	"errors"
	"sync"
)

// Cache implement storager interface
type Cache struct {
	m map[string]interface{}
	*sync.RWMutex
}

// CacheInit init cache
func CacheInit() (*Cache, error) {
	return &Cache{
		m:       make(map[string]interface{}),
		RWMutex: new(sync.RWMutex),
	}, nil
}

// Put put
func (c *Cache) Put(ctx context.Context, key string, val interface{}, opts ...interface{}) error {
	c.Lock()
	c.m[key] = val
	c.Unlock()
	return nil
}

// Get get
func (c *Cache) Get(ctx context.Context, key string, opts ...interface{}) (interface{}, error) {
	c.RLock()
	defer c.RUnlock()
	if v, ok := c.m[key]; ok {
		return v, nil
	}
	return nil, errors.New("invalid value")
}

// Delete delete
func (c *Cache) Delete(ctx context.Context, key string, opts ...interface{}) error {
	delete(c.m, key)
	return nil
}

// Do do
func (c *Cache) Do(ctx context.Context, op interface{}) (interface{}, error) {
	return nil, nil
}

// Close close
func (c *Cache) Close() error {
	return nil
}
