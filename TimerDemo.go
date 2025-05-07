package main

import (
	"fmt"
	"time"
)

func main() {
	timerA := time.NewTimer(2 * time.Second)

	<-timerA.C
	fmt.Println("Timer A expired")

	timerB := time.NewTimer(1 * time.Second)
	go func() {
		<-timerB.C
		fmt.Println("Timer B expired")
	}()
	stopB := timerB.Stop()
	if stopB {
		fmt.Println("Timer B stopped")
	}

}
