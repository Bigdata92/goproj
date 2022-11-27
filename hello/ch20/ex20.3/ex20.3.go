package main

import (
	"goproj/hello/hello/ch20/koreaPost"

	"github.com/tuckersGo/musthaveGo/tree/master/ch20/fedex"
	"github.com/tuckersGo/musthaveGo/tree/master/ch20/koreaPost"
)

func SendBook(name string, sender *fedex.FedexSender) {
	sender.Send(name)
}

func main() {
	sender := &koreaPost.PostSender{}
	SendBook("어린 왕자", sender)
	SendBook("그리스인 조르바", sender)
}
