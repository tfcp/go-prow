package demo

import (
	"prow/internal/service/demo"
	"prow/internal/test"
	"github.com/gogf/gf/test/gtest"
	"testing"
)

var (
	svc *demo.HelloService
)

func init() {
	test.SetupServer()
	svc = demo.NewHelloService()
}

func Test_HelloService(t *testing.T) {
	// one
	gtest.C(t, func(t *gtest.T) {
		whereCondition := map[string]interface{}{
			"name": "tom",
		}
		_, err := svc.One(whereCondition)
		t.Assert(err, nil)
	})
	// list
	gtest.C(t, func(t *gtest.T) {
		whereCondition := map[string]interface{}{
			"name": "",
		}
		_, err := svc.List(whereCondition)
		t.Assert(err, nil)
	})
}
