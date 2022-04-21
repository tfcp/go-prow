package ci

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"prow/internal/model"
)

type Owner struct {
	*model.Model
	Name     string `json:"name"`
	Path     string `json:"path"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Type     string `json:"type"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

func (this *Owner) TableName() string{
	tableName := "owner"
	return fmt.Sprintf("%s%s",g.Config().GetString("database.ci.prefix"),tableName)
}

func (this *Owner) GetOwners(where map[string]interface{}) ([]*Owner, error) {
	var owners []*Owner
	err := model.Db.Where(where).Find(&owners).Error
	if err != nil {
		return owners, err
	}
	return owners, nil
}
