package main

import (
	"fmt"
	"time"
)

func whereFrom(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	whereFrom("direct")

	go whereFrom("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("direct done")
}
