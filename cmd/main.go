package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/luopengift/log"
	"github.com/luopengift/storage"
)

func main() {
	var wg sync.WaitGroup
	store, err := storage.CacheInit()
	if err != nil {
		log.Info("%v", err)
	}
	ctx := context.Background()
	store.Put(ctx, "ttt", "boo")
	fmt.Println(store)
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			err := store.Put(ctx, "ttt", i)
			if err != nil {
				log.Error("%v", err)
			}
			//fmt.Println(tt)
		}(i)
	}
	wg.Wait()
}
