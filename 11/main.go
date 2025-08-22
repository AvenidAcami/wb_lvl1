package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}

	res := make([]int, 0)

	aMap := make(map[int]struct{})
	for _, val := range a {
		aMap[val] = struct{}{}
	}

	for _, val := range b {
		if _, ok := aMap[val]; ok {
			res = append(res, val)
		}
	}

	fmt.Println(res)
}
