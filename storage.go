package storage

import "context"

// Connecter interface
type Connecter interface {
	Open() Storager
}

// Configer config interface
type Configer interface {
	Parse(interface{}) error
}

// Storager is a commen store interface
type Storager interface {
	Put(ctx context.Context, key string, value map[string]interface{}, opts ...interface{}) error
	Get(ctx context.Context, key string, value map[string]interface{}, opts ...interface{}) (interface{}, error)
	Delete(ctx context.Context, key string, value map[string]interface{}, opts ...interface{}) error
	Do(ctx context.Context, op interface{}, opts ...interface{}) (interface{}, error)
	Close() error
}

// New Storager
func New(name, dsn string) Storager {
	return nil
}
