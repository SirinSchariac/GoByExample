package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	//必须要传指针，不然不能实际实现Done操作，导致死锁
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	wg.Done()

	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var waitG sync.WaitGroup

	for i := 1; i <= 5; i++ {
		waitG.Add(1)

		go func() {
			worker(i, &waitG)
		}()
	}

	waitG.Wait()

}
