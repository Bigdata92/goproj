package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int) // create channel

	wg.Add(1)
	go square(&wg, ch) // create goroutine
	ch <- 9            // chan에 data 넣음
	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {
	n := <-ch // data 빼옴

	time.Sleep(time.Second)
	fmt.Printf("Square: %d\n", n*n)
	wg.Done()
}
