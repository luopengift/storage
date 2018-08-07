package storage

import (
	"context"
	"fmt"
	"testing"
)

func TestCache(t *testing.T) {
	store, err := CacheInit()
	if err != nil {
		t.Error(err)
	}
	ctx := context.Background()
	store.Put(ctx, "ttt", "boo")
	tt, err := store.Get(ctx, "ttt")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(tt)

}
