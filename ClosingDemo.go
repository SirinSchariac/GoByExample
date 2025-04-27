package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("Received job:", j)
			} else {
				fmt.Println("No more jobs")
				done <- true
				return
			}
		}
	}()

	for j := 0; j < 4; j++ {
		jobs <- j
		fmt.Println("Sent job:", j)
	}
	close(jobs)
	fmt.Println("Sent all jobs")
	<-done

	_, ok := <-jobs
	fmt.Println("Received more jobs:", ok)
}
