package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "BufferA"
	messages <- "BufferB"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
