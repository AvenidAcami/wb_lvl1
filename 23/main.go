package main

import "fmt"

func main() {
	nums := []int{1, 5, 7, 8, 9, 10}
	words := []string{"some", "word"}
	fmt.Println(delInd(nums, 2))
	fmt.Println(delInd(words, 0))
	words = []string{"some", "word"}
	fmt.Println(delInd(words, 3))
}

func delInd[T any](vals []T, ind int) ([]T, error) {
	if (ind < 0) || (ind >= len(vals)) {
		return vals, fmt.Errorf("index outside array bounds")
	}
	var zero T
	vals[len(vals)-1] = zero
	copy(vals[ind:], vals[ind+1:])
	return vals[:len(vals)-1], nil
}
