package main

import "fmt"

func main() {
	temp := 33

	if temp > 28 {
		fmt.Println("에어컨을 킨다.")
	} else if temp <= 3 {
		fmt.Println("히터를 킨다.")
	} else {
		fmt.Println("대기한다.")
	}
}
