package model

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"
	"time"
)

var (
	Db *gorm.DB
	DbTest *gorm.DB
)

type Model struct {
	ID       int    `gorm:"primary_key" json:"id"`
	CreateAt string `json:"create_at"`
}

func (this *Model) GetOffset(page, size int) int {
	if page < 1 {
		page = 1
	}
	return (page - 1) * size
}

// common where condition
func (this *Model) CommonWhere(where map[string]interface{}) *gorm.DB {
	db := Db
	// like search
	if _, ok := where["name"]; ok {
		db = db.Where("name like ?", "%"+where["name"].(string)+"%")
		delete(where, "name")
	}
	return db
}

// filter form (dont let 0 or "" business)
func (this *Model) FilterBlankForm() *Model {
	return this
}

func Setup() error {
	var err error
	Db, err = setDb(Db, "demo")
	if err != nil {
		return errors.New(fmt.Sprintf("models.Setup err: %v", err))
	}
	DbTest, err = setDb(DbTest, "test")
	if err != nil {
		return errors.New(fmt.Sprintf("models.TestSetup err: %v", err))
	}
	return nil
}

func setDb(db *gorm.DB, dbName string) (*gorm.DB, error) {
	var err error
	db, err = gorm.Open(g.Config().GetString("database."+dbName+".type"), fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		g.Config().GetString("database."+dbName+".user"),
		g.Config().GetString("database."+dbName+".pass"),
		g.Config().GetString("database."+dbName+".host"),
		g.Config().GetString("database."+dbName+".name")))
	if err != nil {
		return nil, err
	}
	db.LogMode(g.Config().GetBool("database." + dbName + ".log"))
	// 多数据库不能用这个属性
	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	//return g.Config().GetString("database."+dbName+".prefix") + defaultTableName
	//	return g.Config().GetString("database.prefix") + defaultTableName
	//}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Hour)
	//Db.Callback().Create().After("gorm:create_at", updateTimeStampForCreateCallback)
	//Db.Callback().Create().After("gorm:create_at", updateTimeStampForCreateCallback)
	//Db.Callback().Update().After("updateTimeStampForCreateCallback")
	fmt.Println("db init success: ", dbName)
	return db, nil
}

// 注册新建钩子在持久化之前
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Format("2006-01-01 03:04:00")
		if createTimeField, ok := scope.FieldByName("CreateAt"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}
	}
}

// 注册更新钩子在持久化之前
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_at"); !ok {
		scope.SetColumn("UpdatedTime", time.Now().Unix())
	}
}
