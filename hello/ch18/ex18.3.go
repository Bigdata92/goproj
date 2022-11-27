package main

import "fmt"

func main() {

	var slice []int

	for i := 1; i <= 10; i++ {
		slice = append(slice, i) // 요소 한개 추가
	}

	slice = append(slice, 11, 12, 13, 14, 15) // 요소 여러개 추가
	fmt.Println(slice)
}
