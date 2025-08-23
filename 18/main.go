package main

import (
	"fmt"
	"sync"
)

type counter struct {
	val int
	mu  sync.Mutex
}

func main() {
	var c counter
	var wg sync.WaitGroup

	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go c.inc(&wg)
	}
	wg.Wait()
	fmt.Println(c.val)

}

func (c *counter) inc(wg *sync.WaitGroup) {
	defer wg.Done()
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val++
}
