package storage

import (
	"context"
	"fmt"
	"sync"
	"testing"
)

func TestCache(t *testing.T) {
	var wg sync.WaitGroup
	store, err := CacheInit()
	if err != nil {
		t.Error(err)
	}
	ctx := context.Background()
	store.Put(ctx, "ttt", "boo")
	fmt.Println(store)
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := store.Get(ctx, "ttt")
			if err != nil {
				t.Error(err)
			}
			//fmt.Println(tt)
		}()
	}
	wg.Wait()
}
