package main

import "fmt"

func main() {
	temps := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, -20, -29.9, 10, 19.9}
	groupedTemps := make(map[int][]float32)

	for _, val := range temps {
		key := int(val) / 10 * 10
		groupedTemps[key] = append(groupedTemps[key], val)
	}

	for key, val := range groupedTemps {
		fmt.Printf("%d: %v\n", key, val)
	}
}
