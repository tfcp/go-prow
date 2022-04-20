package ormgen

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
)

var (
	// 生成model的路径(可以自定义 但是需要确保路径文件存在)
	path = "./tools/ormgen/model/temp.go"
	// mysql地址
	//dsn = "user:pwd@tcp(localhost:3306)/database?charset=utf8"

)

func BaseTool(tableName string, dbName string) {
	user := g.Config().GetString("database." + dbName + ".user")
	pass := g.Config().GetString("database." + dbName + ".pass")
	host := g.Config().GetString("database." + dbName + ".host")
	port := g.Config().GetString("database." + dbName + ".port")
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		user, pass, host, port, dbName,
	)
	t2t := NewTable2Struct()
	// table 需要的table名称()
	err := t2t.
		// table 需要转换的table名称(不指定则默认全部表转换)
		Table(tableName).
		// 包名
		PackageName("model").
		SavePath(path).
		// 是否需要json标签
		EnableJsonTag(true).
		// 是否生成对应的method方法
		MakeOrmMethod(true).
		Dsn(dsn).
		Run()
	if err != nil {
		fmt.Println(err)
	}
}
