package storage

import (
	"context"
	"fmt"
	"testing"
)

func TestFile(t *testing.T) {
	f, err := FileInit("/tmp/test")
	if err != nil {
		t.Error(err)
	}
	ctx := context.Background()
	a, err := f.Put(ctx, "a", "b")
	if err != nil {
		t.Error(err)
	}
	a, err = f.Get(ctx, "a")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(a)
}
