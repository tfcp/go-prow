package consumer

import (
	"fmt"
	"time"
)

var (
	// notice que data userId
	noticeQ = make(chan []string, 100)
)

func SetMergeNotice(data []string) {
	noticeQ <- data
}

// watch notice data
func RunMergeNotice() {
	fmt.Println("mergeNotice consumer init success")
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
