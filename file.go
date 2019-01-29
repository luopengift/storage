package storage

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
)

//var _ Storager = &File{}

// File implement storager interface
type File struct {
	BaseDir string
}

// FileInit file init
func FileInit(dir string) (*File, error) {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}
	return &File{
		BaseDir: dir,
	}, nil
}

// Put xx
func (f *File) Put(ctx context.Context, key, val string, opts ...interface{}) error {
	return ioutil.WriteFile(filepath.Join(f.BaseDir, key), []byte(val), 0644)
}

// Get xx
func (f *File) Get(ctx context.Context, key string, val interface{}, opts ...interface{}) ([]byte, error) {
	return ioutil.ReadFile(filepath.Join(f.BaseDir, key))
}

// Delete xx
func (f *File) Delete(ctx context.Context, key string, res map[string]interface{}, opts ...interface{}) error {
	return os.Remove(filepath.Join(f.BaseDir, key))
}

// Do x
func (f *File) Do(ctx context.Context, op interface{}, opts ...interface{}) (interface{}, error) {
	return nil, nil
}

// Close xx
func (f *File) Close() error {
	return nil
}
