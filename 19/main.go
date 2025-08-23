package main

import "fmt"

func main() {
	s := "главрыба"
	fmt.Println(reverse(s))
}

func reverse(s string) string {
	str := []rune(s)
	lenStr := len(str)

	for i := 0; i < len(str)/2; i++ {
		str[i], str[lenStr-1-i] = str[lenStr-1-i], str[i]
	}
	return string(str)
}
