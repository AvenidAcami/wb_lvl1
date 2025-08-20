package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var workersCount int
	fmt.Scan(&workersCount)

	ch := make(chan int)

	for i := 1; i <= workersCount; i++ {
		go worker(ch)
	}

	for {
		ch <- rand.Intn(100)
	}

}

func worker(ch chan int) {
	for j := range ch {
		fmt.Println(j)
	}
}
