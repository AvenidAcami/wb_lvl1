package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	var workersCount int
	var wg sync.WaitGroup

	fmt.Scan(&workersCount)

	ch := make(chan int)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)

	wg.Add(workersCount)

	for i := 1; i <= workersCount; i++ {
		go worker(ch, i, &wg)
	}

	go func() {
		for {
			select {
			case <-sigChan:
				close(ch)
				return
			default:
				ch <- rand.Intn(100)
			}
		}
	}()

	wg.Wait()
	fmt.Println("main stopped")
}

func worker(ch chan int, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range ch {
		fmt.Println(j, "worker", id)
	}
	fmt.Println("worker", id, "stopped")
}
