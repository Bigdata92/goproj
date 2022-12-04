package main

import (
	"fmt"
)

func main() {
	ch := make(chan int) // create size 0 channel

	ch <- 9
	fmt.Println("Never print")

}
