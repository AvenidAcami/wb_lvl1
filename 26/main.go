package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isOnlyUniqueSymbs("abcd"))
	fmt.Println(isOnlyUniqueSymbs("abCdefAaf"))
	fmt.Println(isOnlyUniqueSymbs("aabcd"))
}

func isOnlyUniqueSymbs(s string) bool {
	var lowerSymb string
	symbs := make(map[string]struct{})

	for _, val := range s {
		lowerSymb = strings.ToLower(string(val))
		if _, ok := symbs[lowerSymb]; ok {
			return false
		} else {
			symbs[lowerSymb] = struct{}{}
		}
	}
	return true
}
