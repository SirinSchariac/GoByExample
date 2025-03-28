package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func tripleAdd(a int, b int, c int) int {
	return a + b + c
}

func main() {
	res := add(1, 2)
	fmt.Println("1+2 =", res)

	res = tripleAdd(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
