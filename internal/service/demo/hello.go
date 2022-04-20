package demo

import (
	"prow/internal/model/demo"
	"prow/library/log"
)

type HelloService struct {
	helloModel *demo.Hello
}

func NewHelloService() (s *HelloService) {
	h := &HelloService{}
	h.helloModel = &demo.Hello{}
	return h
}

func (this *HelloService) List(where map[string]interface{}) ([]*demo.Hello, error) {
	list, err := this.helloModel.ListHello(where)
	if err != nil {
		log.Logger.Errorf("HelloService ListError: %v", err)
		return nil, err
	}
	return list, nil
}

func (this *HelloService) One(where map[string]interface{}) (*demo.Hello, error) {
	one, err := this.helloModel.OneHello(where)
	if err != nil {
		log.Logger.Errorf("HelloService OneError: %v", err)
		return nil, err
	}
	return &one, nil
}
