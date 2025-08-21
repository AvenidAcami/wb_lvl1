package main

import (
	"math/rand/v2"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	nums := make(map[int]int)

	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()

			mu.Lock()
			defer mu.Unlock()

			nums[rand.IntN(10000)] = rand.IntN(10000)

		}()
	}
	wg.Wait()
}
