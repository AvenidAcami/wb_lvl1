package main

import "fmt"

func main() {
	a := 123
	b := 53

	a = a + b
	b = a - b
	a = a - b

	fmt.Println(a, b)
}
