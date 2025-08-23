package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 5, 7, 8, 9, 10}
	for _, val := range nums {
		fmt.Println(binSearch(nums, val))
	}
	fmt.Println(binSearch(nums, 0))
	fmt.Println(binSearch(nums, 11))
}

func binSearch(nums []int, val int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		middle := (right + left) / 2
		if nums[middle] == val {
			return middle
		}
		if nums[middle] < val {
			left = middle + 1
		} else {
			right = middle - 1
		}

	}
	return -1
}
