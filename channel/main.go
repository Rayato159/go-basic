package main

import (
	"fmt"
	"time"
)

func main() {
	jobCh := make(chan int, 10)
	resultCh := make(chan int, 10)

	for i := range 10 {
		jobCh <- i + 1
	}
	close(jobCh)

	for range 2 {
		go double(jobCh, resultCh)
	}

	for range 10 {
		result := <-resultCh
		fmt.Println(result)
	}
}

func double(jobCh <-chan int, resultCh chan<- int) {
	for j := range jobCh {
		time.Sleep(1 * time.Second)
		resultCh <- j * 2
	}
}
