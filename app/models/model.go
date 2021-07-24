package models

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/jinzhu/gorm"
	"log"
	"reflect"
	"time"
)

var gormDb *gorm.DB

var (
	tablePrefix = "prow_"
	Config      = config.ReadFromJson("./config/config.json")
)

type Model struct {
	ID       int       `gorm:"primary_key" json:"id"`
	CreateAt time.Time `json:"create_at" zh:"创建时间" en:"Create Time"`
	UpdateAt time.Time `json:"update_at" zh:"更新时间" en:"Update Time"`
}

var Generators = map[string]table.Generator{
	"owner":   GetProwOwnerTable,
	"notice":  GetProwNoticeTable,
	"path":    GetProwPathTable,
	"project": GetProwProjectTable,
	"robot":   GetProwRobotTable,

	// generators end
}

//// Setup initializes the database instance
func Setup() {
	var err error
	dbConfig := Config.Databases
	gormDb, err = gorm.Open(config.DriverMysql, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.GetDefault().User,
		dbConfig.GetDefault().Pwd,
		dbConfig.GetDefault().Host,
		dbConfig.GetDefault().Name))
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	gormDb.SingularTable(true)
	gormDb.DB().SetMaxIdleConns(10)
	gormDb.DB().SetMaxOpenConns(100)
	log.Println("db init success")
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer gormDb.Close()
}

// addExtraSpaceIfExist adds a separator
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}

func Lan(model interface{}, field string) string {
	t := reflect.TypeOf(model).Elem()
	doc := make(map[string]string)
	lan := Config.Language
	if lan != "en" && lan != "zh" {
		lan = "zh"
	}
	for i := 0; i < t.NumField(); i++ {
		doc[t.Field(i).Tag.Get("json")] = t.Field(i).Tag.Get(lan)
	}
	return doc[field]
}
