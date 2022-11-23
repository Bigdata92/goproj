package main

import "fmt"

const PI = 3.14 // type x
const FloatPI float64 = 3.14

func main() {
	var a int = PI * 100
	// var b int = FloatPI * 100	// type error

	fmt.Println(a)
}
