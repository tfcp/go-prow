package demo

import (
	"fmt"
	"time"
)

func HelloProcess() {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("this is demo-hello process")
	}
}

func TestProcess() {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("this is demo-test process")
	}
}
