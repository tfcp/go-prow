package service

import "fmt"

func Notice(userId int) {
	fmt.Println("consumer notice:", userId)
}
