package main

import "fmt"

func main() {
	var minimumWage int = 10
	var workingHour int = 20

	// 변수 선언 및 초기화
	var income int = minimumWage * workingHour

	fmt.Println(minimumWage, workingHour, income)
}
