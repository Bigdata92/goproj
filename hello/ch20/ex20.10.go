package main

type Reader interface {
	Read()
}

type Closer interface {
	Close()
}

type File struct {
}

func (f *File) Read() {
}

func ReadFile(reader Reader) {
	// 런타임 에러가 발생합니다.
	c := reader.(Closer)
	c.Close()
}

func main() {
	file := &File{}
	ReadFile(file)
}
