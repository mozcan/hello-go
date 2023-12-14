package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	width, height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(geo geometry) {
	fmt.Println(geo.perim())
	fmt.Println(geo.area())
}

func main() {
	rect := rectangle{3.11, 3.12}
	crcl := circle{3}

	measure(rect)
	measure(crcl)
}
