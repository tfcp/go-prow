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
			// merge notice
			fmt.Println("this is mr notice:", nt)
			// robot git merge reply
			fmt.Println("this is mr robot comment:", nt)
		default:

		}
		time.Sleep(1 * time.Second)
	}
}
