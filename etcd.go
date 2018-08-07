package storage

import (
	"context"
	"time"

	"github.com/coreos/etcd/clientv3"
)

// Etcd implement strorager interface
type Etcd struct {
	*clientv3.Client
}

// EtcdInit storager
func EtcdInit(endpoints []string, timeout int) (Storager, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: time.Duration(timeout) * time.Second,
	})
	return &Etcd{cli}, err
}

// Interface interface
func (etcd *Etcd) Interface() interface{} {
	return etcd.Client
}

// Put xx
func (etcd *Etcd) Put(ctx context.Context, key, val string, opts ...interface{}) (interface{}, error) {
	return etcd.Client.Put(ctx, key, val)
}

// Get xx
func (etcd *Etcd) Get(ctx context.Context, key string, opts ...interface{}) (interface{}, error) {
	return etcd.Client.Get(ctx, key)
}

// Delete xx
func (etcd *Etcd) Delete(ctx context.Context, key string, opts ...interface{}) (interface{}, error) {
	return etcd.Client.Delete(ctx, key)
}

// Do xx
func (etcd *Etcd) Do(ctx context.Context, op interface{}) (interface{}, error) {
	return nil, nil //etcd.Client.Do(ctx, nil)
}

// Close xx
func (etcd *Etcd) Close() error {
	return etcd.Client.Close()
}
