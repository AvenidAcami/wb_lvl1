package main

import "fmt"

func main() {
	nums := []int{10, 7, 8, 9, 1, 5}
	fmt.Println(quickSort(nums))
}

func quickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	pivot := nums[0]
	left := make([]int, 0)
	middle := make([]int, 0)
	right := make([]int, 0)

	for _, val := range nums {
		if val == pivot {
			middle = append(middle, val)
		}
		if val < pivot {
			left = append(left, val)
		}
		if val > pivot {
			right = append(right, val)
		}
	}
	qLeftAndMiddle := append(quickSort(left), middle...)
	return append(qLeftAndMiddle, quickSort(right)...)
}
