package main

import "fmt"

func zeroVal(val int) {
	val = 0
}

func zeroPtr(valPtr *int) {
	*valPtr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroVal(i)
	fmt.Println("zeroVal:", i)

	zeroPtr(&i)
	fmt.Println("zeroPtr:", i)

	fmt.Println("pointer:", &i)
}
