package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	nums := make(chan int)
	timer := time.After(3 * time.Second)

	go func() {
		for {
			select {
			case <-timer:
				close(nums)
				return
			case nums <- rand.IntN(100):
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	for j := range nums {
		fmt.Println(j)
	}
}
