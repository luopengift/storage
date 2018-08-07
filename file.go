package storage

import (
	"context"
	"io/ioutil"
	"os"
	"path"
)

// File implement storager interface
type File struct {
	BaseDir string
}

// FileInit file init
func FileInit(dir string) (Storager, error) {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}
	return &File{
		BaseDir: dir,
	}, nil
}

// Interface interface
func (f *File) Interface() interface{} {
	return nil
}

// Put xx
func (f *File) Put(ctx context.Context, key, val string, opts ...interface{}) (interface{}, error) {
	err := ioutil.WriteFile(path.Join(f.BaseDir, key), []byte(val), 0644)
	return nil, err
}

// Get xx
func (f *File) Get(ctx context.Context, key string, opts ...interface{}) (interface{}, error) {
	return ioutil.ReadFile(path.Join(f.BaseDir, key))
}

// Delete xx
func (f *File) Delete(ctx context.Context, key string, opts ...interface{}) (interface{}, error) {
	return nil, os.Remove(path.Join(f.BaseDir, key))
}

// Do x
func (f *File) Do(ctx context.Context, op interface{}) (interface{}, error) {
	return nil, nil
}

// Close xx
func (f *File) Close() error {
	return nil
}
