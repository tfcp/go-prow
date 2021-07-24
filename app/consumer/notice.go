package consumer

import (
	"fmt"
	"time"
)

var (
	// notice que data userId
	noticeQ = make(chan int, 100)
)

func SetNotice(data int) {
	noticeQ <- data
}

// watch notice data
func RunNotice() {
	fmt.Println("notice consumer init success")
	// todo 日志
	for {
		select {
		case nt := <-noticeQ:
			fmt.Println("this is nt:", nt)
		default:

		}
		time.Sleep(1 * time.Second)
	}
}
