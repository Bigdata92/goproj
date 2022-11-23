package main

import "fmt"

func main() {
	const C int = 10 // 상수 선언

	var b int = C * 20
	C = 20
	fmt.Println(&C) // 상수 mem 주소 접근 불가
}
