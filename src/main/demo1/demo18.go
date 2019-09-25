package main

import (
	"fmt"
	"math"
)


type geometry interface {
	area() float64
	perim() float64
}

type jiheti struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (j jiheti) area() float64 {
	return j.width * j.height
}

func (j jiheti) perim() float64 {
	return 2*j.width + 2*j.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	j := jiheti{width: 3, height: 4}
	c := circle{radius: 5}

	measure(j)
	measure(c)
}