package main

import (
	"fmt"
	"sync"
)

func main() {
	// fmt.Println("Hello main thread 1")
	// time.Sleep(1 * time.Second)

	// go hello()

	// fmt.Println("Hello main thread 2")

	var wg sync.WaitGroup
	wg.Add(3)

	for i := range 3 {
		go func(i int, wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Println(i)
		}(i, &wg)
	}

	wg.Wait()
}

// func hello() {
// 	for {
// 		fmt.Println("Hello from another thread")
// 	}
// }
