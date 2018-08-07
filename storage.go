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
	Interface() interface{}
	Put(ctx context.Context, key, val string, opts ...interface{}) (interface{}, error)
	Get(ctx context.Context, key string, opts ...interface{}) (interface{}, error)
	Delete(ctx context.Context, key string, opts ...interface{}) (interface{}, error)
	Do(ctx context.Context, op interface{}) (interface{}, error)
	Close() error
}
