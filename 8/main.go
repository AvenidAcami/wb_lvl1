package main

import (
	"fmt"
)

func main() {
	var n, i int64
	var val bool
	fmt.Scan(&n)
	fmt.Scan(&i)
	if (i < 1) || (i > 64) {
		fmt.Println("Некорректная степень")
		return
	}
	fmt.Scan(&val)

	if val {
		fmt.Println(n | (1 << (i - 1)))
	} else {
		fmt.Println(n &^ (1 << (i - 1)))
	}
}
