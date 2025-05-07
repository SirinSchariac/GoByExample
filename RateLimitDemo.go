package main

import (
	"fmt"
	"time"
)

const (
	MAX_REQUEST = 3
	RATE_LIMIT  = 400 * time.Millisecond
)

func main() {
	req := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		req <- i
	}
	close(req)

	limiter := time.Tick(RATE_LIMIT)

	for r := range req {
		<-limiter
		fmt.Println("Request", r, time.Now())
	}

	//bursty limiter with request buffers
	burstyLimiter := make(chan time.Time, MAX_REQUEST)

	for i := 0; i < MAX_REQUEST; i++ {
		burstyLimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(RATE_LIMIT) {
			burstyLimiter <- t
		}
	}()
	burstyReq := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyReq <- i
	}
	close(burstyReq)
	for r := range burstyReq {
		<-burstyLimiter
		fmt.Println("Bursty Request", r, time.Now())
	}
}
