package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
}

func (s *Student) String() string {
	return "Student"
}

type Actor struct {
}

func (a *Actor) String() string {
	return "Actor"
}

func ConvertType(stringer Stringer) {
	// stringer값이 *Student 타입이 아니기 때문에 에러가 발생합니다.
	student := stringer.(*Student)
	fmt.Println(student)
}

func main() {
	actor := &Actor{}
	ConvertType(actor)
}
