package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(whatType(123))
	fmt.Println(whatType("fadsaa"))
	fmt.Println(whatType(true))
	fmt.Println(whatType(make(chan int)))
}

func whatType(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	default:
		if reflect.TypeOf(v).Kind() == reflect.Chan {
			return "chan"
		}
	}
	return "unknown"
}
