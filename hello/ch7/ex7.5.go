package main

import "fmt"

func Divide(a, b int) (result int, success bool) { // 반환 변수명 명시
	if b == 0 {
		result = 0
		success = false
		return // 명시적으로 반환값 지정 x return 문
	}
	result = a / b
	success = true
	return
}

func main() {
	c, success := Divide(9, 3)
	fmt.Println(c, success)
	d, success := Divide(9, 0)
	fmt.Println(d, success)
}
