package main

import (
	"fmt"
	"sort"
)

func main() {

	s := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(s) // 정렬
	fmt.Println(s)
}
