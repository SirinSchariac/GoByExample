package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("The area is ", g.area())
	fmt.Println("The perim is ", g.perim())
}

func detectCircle(g geometry) {
	if c, ok := g.(geometry).(circle); ok {
		fmt.Println("This is a circle with radius:", c.radius)
	} else {
		fmt.Println("This is not a circle")
	}
}

func main() {
	r := rect{width: 10, height: 5}
	c := circle{radius: 5}

	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}
