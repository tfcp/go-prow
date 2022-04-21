package demo

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"prow/internal/model"
)


type Hello struct {
	*model.Model
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (this *Hello) TableName() string{
	tableName := "hello"
	return fmt.Sprintf("%s%s",g.Config().GetString("database.demo.prefix"),tableName)
}

func (this *Hello) ListHello(where map[string]interface{}) ([]*Hello, error) {
	var hellos []*Hello
	err := model.Db.Where(where).Find(&hellos).Error
	if err != nil {
		return hellos, err
	}
	return hellos, nil
}

func (this *Hello) OneHello(where map[string]interface{}) (Hello, error) {
	var hello Hello
	err := model.Db.Where(where).First(&hello).Error
	if err != nil {
		return hello, err
	}
	return hello, nil
}
