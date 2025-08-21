package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	// Выход по условию
	go func() {
		c := 0
		for {
			c++
			time.Sleep(200 * time.Millisecond)
			if c == 5 {
				fmt.Println("goroutine", 1, "stopped")
				return
			}
		}
	}()

	// Выход через канал уведомления
	ch := make(chan byte, 1)

	go func() {
		<-ch
		fmt.Println("goroutine", 2, "stopped")
	}()
	time.Sleep(1 * time.Second)
	ch <- 1
	close(ch)

	// Выход через закрытие канала
	ch1 := make(chan byte, 1)
	go func() {
		for range ch1 {
		}
		fmt.Println("goroutine", 3, "stopped")
	}()
	time.Sleep(1 * time.Second)
	close(ch1)

	// Выход через контекст
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-ctx.Done()
		fmt.Println("goroutine", 4, "stopped")
	}()
	time.Sleep(1 * time.Second)
	cancel()

	// Выход через Goexit
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("goroutine", 5, "stopped")
		runtime.Goexit()
	}()
	time.Sleep(2 * time.Second)
}
