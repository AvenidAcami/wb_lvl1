package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func (p *Point) Distance(other Point) float64 {
	xDiff := p.x - other.x
	yDiff := p.y - other.y

	return math.Sqrt((yDiff*yDiff + xDiff*xDiff))
}

func main() {
	point1 := NewPoint(10, -3.1)
	point2 := NewPoint(-1.25, 3.6)

	fmt.Println(point1.Distance(point2))
}
