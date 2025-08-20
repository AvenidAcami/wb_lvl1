package main

import "fmt"

type human struct {
	name string
}

func (hm human) getName() string {
	return hm.name
}

type action struct {
	human
}

func main() {
	action := action{
		human{
			name: "Anton",
		},
	}
	fmt.Println(action.getName())
}
