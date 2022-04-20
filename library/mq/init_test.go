package mq

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
)

func init() {
	g.Config().SetPath("../../config")
	fmt.Println("mq init...")
}
