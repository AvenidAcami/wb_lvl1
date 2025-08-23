package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	go func() {
		fmt.Println("some work")
		sleep(2 * time.Second)
	}()
	sleep(3 * time.Second)
	fmt.Println("end")
}

func sleep(d time.Duration) {
	done := make(chan struct{})
	time.AfterFunc(d, func() {
		done <- struct{}{}
	})
	<-done
}
