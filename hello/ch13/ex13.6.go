package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Age   int32   // 4byte
	Score float64 // 8byte
}

func main() {
	user := User{23, 77.2}
	fmt.Println(unsafe.Sizeof(user))
}
