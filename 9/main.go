package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	sigChan := make(chan os.Signal, 1)
	defer close(sigChan)
	signal.Notify(sigChan, syscall.SIGINT)

	go func() {
		for {
			select {
			case ch1 <- rand.IntN(100):
				time.Sleep(200 * time.Millisecond)
			case <-sigChan:
				close(ch1)
				return
			}

		}
	}()

	go func() {
		for j := range ch1 {
			ch2 <- j * 2
		}
		close(ch2)
	}()

	for j := range ch2 {
		fmt.Println(j)
	}
}
