package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	str1 := "Hello World!"
	str2 := []byte(str1)

	stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	stringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))

	fmt.Printf("str1:\t%x\n", stringHeader1.Data)
	fmt.Printf("str2:\t%x\n", stringHeader2.Data)
}
