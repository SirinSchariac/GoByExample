package main

import "fmt"

func Ping(pings chan<- string, msg string) {
	pings <- msg //chan<- : receive only
}

func Pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings //<-chan : send only
	pongs <- msg   //chan<- : receive only
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	Ping(pings, "ping message\n")
	Pong(pings, pongs)

	fmt.Print(<-pongs)
}
