package main

import "fmt"

type eachElem struct {
	ind int
	val int
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	resChan := make(chan eachElem, len(nums))

	for i, num := range nums {
		go func(ind, num int) {
			resChan <- eachElem{ind, num * num}
		}(i, num)
	}

	for i := 0; i < len(nums); i++ {
		currElem := <-resChan
		nums[currElem.ind] = currElem.val
	}

	fmt.Println(nums)
}
