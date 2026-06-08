package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var x int64
	var wg sync.WaitGroup
	//var mu sync.Mutex
	for range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				//mu.Lock()
				//x++
				//mu.Unlock()
				atomic.AddInt64(&x, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println(x)
}
