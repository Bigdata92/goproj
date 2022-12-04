package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int) {
	for n := range ch {
		fmt.Printf("Square: %d\n", n*n)
		time.Sleep(time.Second)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int) // create channel

	wg.Add(1)
	go square(&wg, ch) // create goroutine

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}

	wg.Wait()
}
