package demo

import (
	"fmt"
	"prow/internal/model"
	"github.com/gogf/gf/frame/g"
	"github.com/jinzhu/gorm"
)

// gorm文档: https://www.tizi365.com/archives/22.html

type User struct {
	*model.Model
	Name         string `json:"name"`
	Age          int    `json:"age"`
	Sex          int    `json:"sex"`
	Status       int    `json:"status"`
	Role         int    `json:"role"`
	Pwd          string `json:"pwd"`
	Avatar       string `json:"avatar"`
	Introduction string `json:"introduction"`
}

func (this *User) TableName() string{
	// 静态表名 动态表名走db.Table
	tableName := "user"
	return fmt.Sprintf("%s%s",g.Config().GetString("database.demo.prefix"),tableName)
}

func (this *User) ListUser(where map[string]interface{}, page, size int) ([]*User, error) {
	var users []*User
	session := this.CommonWhere(where)
	err := session.Where(where).Offset(this.GetOffset(page, size)).Limit(size).Find(&users).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return users, nil
		}
		return users, err
	}
	return users, nil
}

func (this *User) CountUser(where map[string]interface{}) (int, error) {
	var count int
	session := this.CommonWhere(where)
	err := session.Model(this).Where(where).Count(&count).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return count, nil
		}
		return count, err
	}
	return count, nil
}

func (this *User) OneUser(where map[string]interface{}) (User, error) {
	var user User
	err := model.Db.Where(where).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, nil
		}
		return user, err
	}
	return user, nil
}

func (this *User) DeleteUser(where map[string]interface{}) error {
	var user User
	err := model.Db.Where(where).Delete(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	return nil
}

// ps: 当用结构体更新的时候，当结构体的值是""或者0，false等，就什么也不会更新。
//func (this *User) UpdateUser(user User, update map[string]interface{}) error {
func (this *User) UpdateUser(user User, upMap map[string]interface{}) error {
	// 更新具体值
	//err := model.Db.Model(&user).Update(update).Error
	err := model.Db.Model(&user).Update(upMap).Error
	// 更新模型
	//err := model.Db.Update(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (this *User) CreateUser(user User) error {
	err := model.Db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}
