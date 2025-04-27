package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("Working...")
	time.Sleep(time.Second)
	fmt.Print("done!\n")

	done <- true
}

func main() {
	done := make(chan bool, 1)

	go worker(done)

	<-done // wait for worker to finish
	fmt.Println("Main done")
}
