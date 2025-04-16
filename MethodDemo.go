package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func (r *rect) resize(new_width, new_height int) {
	r.width = new_width
	r.height = new_height
}

func main() {
	block := rect{width: 10, height: 5}
	fmt.Println("area:", block.area())
	fmt.Println("perim:", block.perim())
	block.resize(20, 2)
	fmt.Println("area after resize:", block.area())
	fmt.Println("perim after resize:", block.perim())
}
