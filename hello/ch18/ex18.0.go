package main

import "fmt"

var slice = []int{10, 20, 30}

func main() {
	for _, v := range slice {
		fmt.Println(v)
	}
}
