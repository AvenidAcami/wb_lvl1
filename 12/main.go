package main

import "fmt"

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}

	res := make(map[string]struct{})

	for _, val := range words {
		res[val] = struct{}{}
	}

	for key, _ := range res {
		fmt.Println(key)
	}
}
