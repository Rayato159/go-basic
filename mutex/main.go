package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu      sync.Mutex
	balance int = 1000
)

func main() {
	doneCh := make(chan bool, 3)

	go UpdateBalance(doneCh, 100)
	go UpdateBalance(doneCh, 150)
	go UpdateBalance(doneCh, 200)

	<-doneCh
	<-doneCh
	<-doneCh

	fmt.Println(balance)
}

func UpdateBalance(doneCh chan<- bool, amount int) {
	mu.Lock()

	fmt.Println("Updating balance")
	time.Sleep(1 * time.Second)

	balance -= amount
	doneCh <- true

	mu.Unlock()
	fmt.Println("Balance updated")
}
