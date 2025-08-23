package main

import "fmt"

func main() {
	fmt.Println(reverse("a reallylongwordhereandmore"))
	fmt.Println(reverse("short a veryveryverylongword"))
	fmt.Println(reverse("singlewordthatiswaytoolong"))
	fmt.Println(reverse("a b c d e f g h i j superlongword"))
	fmt.Println(reverse("word1 word2 thisisareallylongword"))
}

func reverse(words string) string {
	b := []byte(words)

	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}

	left := 0
	for i := 0; i <= len(b); i++ {
		if i == len(b) || b[i] == ' ' {
			for j := 0; j < (i-left)/2; j++ {
				b[left+j], b[i-1-j] = b[i-1-j], b[left+j]
			}
			left = i + 1
		}
	}

	return string(b)
}
