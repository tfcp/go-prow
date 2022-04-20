package ormgen

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"os/exec"
	"strings"
)

type Table2Struct struct {
	dsn             string
	savePath        string
	db              *sql.DB
	table           string
	tableCap        string
	prefix          string
	err             error
	realNameMethod  string
	enableJsonTag   bool   // 是否添加json的tag, 默认不添加
	enableOrmMethod bool   // 是否生成method(Gf orm)
	enableGinMethod bool   // 是否生成method(Gin orm)
	packageName     string // 生成struct的包名(默认为空的话, 则取名为: package model)
}

type column struct {
	ColumnName    string
	Type          string
	Nullable      string
	TableName     string
	ColumnComment string
	Tag           string
	Json          string
}

var typeForMysqlToGo = map[string]string{
	"int":                "int32",
	"integer":            "int32",
	"tinyint":            "int32",
	"smallint":           "int32",
	"mediumint":          "int32",
	"bigint":             "int64",
	"int unsigned":       "uint32",
	"integer unsigned":   "uint32",
	"tinyint unsigned":   "uint32",
	"smallint unsigned":  "uint32",
	"mediumint unsigned": "uint32",
	"bigint unsigned":    "uint32",
	"bit":                "int",
	"bool":               "bool",
	"enum":               "string",
	"set":                "string",
	"varchar":            "string",
	"char":               "string",
	"tinytext":           "string",
	"mediumtext":         "string",
	"text":               "string",
	"longtext":           "string",
	"blob":               "string",
	"tinyblob":           "string",
	"mediumblob":         "string",
	"longblob":           "string",
	"date":               "string", // time.Time
	"datetime":           "string", // time.Time
	"timestamp":          "string", // time.Time
	"time":               "string", // time.Time
	"float":              "float64",
	"double":             "float64",
	"decimal":            "float64",
	"binary":             "string",
	"varbinary":          "string",
}

func NewTable2Struct() *Table2Struct {
	return &Table2Struct{}
}

func (t *Table2Struct) Table(tab string) *Table2Struct {
	t.table = tab
	return t
}

func (t *Table2Struct) EnableJsonTag(p bool) *Table2Struct {
	t.enableJsonTag = p
	return t
}

/**
 * 是否生成Gf orm方法
 */
func (t *Table2Struct) MakeOrmMethod(p bool) *Table2Struct {
	t.enableOrmMethod = p
	// 只能有一种生效
	if p == true {
		t.enableGinMethod = false
	}
	return t
}

/**
 * 是否生成Gin orm方法
 */
func (t *Table2Struct) MakeGinMethod(p bool) *Table2Struct {
	t.enableGinMethod = p
	// 只能有一种生效
	if p == true {
		t.enableOrmMethod = false
	}
	return t
}

func (t *Table2Struct) PackageName(r string) *Table2Struct {
	t.packageName = r
	return t
}

func (t *Table2Struct) SavePath(p string) *Table2Struct {
	t.savePath = p
	return t
}

func (t *Table2Struct) Dsn(d string) *Table2Struct {
	t.dsn = d
	return t
}

func (t *Table2Struct) Run() error {
	// 链接mysql, 获取db对象
	t.dialMysql()
	if t.err != nil {
		return t.err
	}

	// 获取表和字段的shcema
	tableColumns, err := t.getColumns()
	if err != nil {
		return err
	}

	//fmt.Println(tableColumns)

	// 包名
	var packageName string
	if t.packageName == "" {
		packageName = "package model\n\n"
	} else {
		packageName = fmt.Sprintf("package %s\n\n", t.packageName)
	}

	// 组装struct
	var structContent string
	for tableRealName, item := range tableColumns {
		// 去除前缀
		if t.prefix != "" {
			tableRealName = tableRealName[len(t.prefix):]
		}
		tableName := tableRealName
		//switch len(tableName) {
		//case 0:
		//case 1:
		//	tableName = strings.ToUpper(tableName[0:1])
		//default:
		//	// 字符长度大于1时
		//	tableName = strings.ToUpper(tableName[0:1]) + tableName[1:]
		//}
		depth := 1
		structContent += "type " + t.camelCase(tableName) + " struct {\n"
		structContent += "	*model.Model\n"

		for _, v := range item {
			//structContent += tab(depth) + v.ColumnName + " " + v.Type + " " + v.Json + "\n"
			// 字段注释
			var clumnComment string
			if v.ColumnComment != "" {
				clumnComment = fmt.Sprintf(" // %s", v.ColumnComment)
			}
			v = FilterToGorm(v)
			if t.enableJsonTag {
				structContent += fmt.Sprintf("%s%s %s %s %s\n",
					tab(depth), v.ColumnName, v.Type, v.Json, clumnComment)
			} else {
				structContent += fmt.Sprintf("%s%s %s  %s\n",
					tab(depth), v.ColumnName, v.Type, clumnComment)
			}
		}
		structContent += tab(depth-1) + "}\n\n"

		// 添加 method 获取真实表名
		if t.realNameMethod != "" {
			structContent += fmt.Sprintf("func (*%s) %s() string {\n",
				tableName, t.realNameMethod)
			structContent += fmt.Sprintf("%sreturn \"%s\"\n",
				tab(depth), tableRealName)
			structContent += "}\n\n"
		}
	}

	// 如果有引入 time.Time, 则需要引入 time 包
	var importContent string
	if strings.Contains(structContent, "time.Time") {
		importContent = "import \"time\"\r\n"
	}
	// 引入import相关
	//if t.enableOrmMethod {
	//	importContent += "import \"database/sql\"\r\n"
	//	importContent += "import \"github.com/gogf/gf/g\"\r\n"
	//	var varString string
	//	varString = "var (\r\n"
	//	varString += "\n" + t.table + " = g.DB(\"" + t.GetDbName() + "\").Table(\"" + t.table + "\").Safe()\r\n"
	//	varString += ")\r\n"
	//	importContent += varString
	//}

	// 写入文件struct
	var savePath = t.savePath
	// 是否指定保存路径
	if savePath == "" {
		savePath = "model.go"
	}
	filePath := fmt.Sprintf("%s", savePath)
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Can not write file")
		return err
	}
	defer f.Close()

	funcString := "\r\n"
	if t.enableOrmMethod {
		funcString = t.makeOrmMethod()
	}

	content := packageName + importContent + structContent + funcString
	f.WriteString(content)

	fmt.Println(content)
	cmd := exec.Command("gofmt", "-w", filePath)
	cmd.Run()

	return nil
}

func tab(depth int) string {
	return strings.Repeat("\t", depth)
}

func (t *Table2Struct) dialMysql() {
	if t.db == nil {
		if t.dsn == "" {
			t.err = errors.New("dsn数据库配置缺失")
			return
		}
		t.db, t.err = sql.Open("mysql", t.dsn)
	}
	return
}

func (t *Table2Struct) getColumns(table ...string) (tableColumns map[string][]column, err error) {
	tableColumns = make(map[string][]column)
	// sql
	var sqlStr = `SELECT COLUMN_NAME,DATA_TYPE,IS_NULLABLE,TABLE_NAME,COLUMN_COMMENT
		FROM information_schema.COLUMNS 
		WHERE table_schema = DATABASE()`
	// 是否指定了具体的table
	if t.table != "" {
		sqlStr += fmt.Sprintf(" AND TABLE_NAME = '%s'", t.prefix+t.table)
	}
	// sql排序
	sqlStr += " order by TABLE_NAME asc, ORDINAL_POSITION asc"

	rows, err := t.db.Query(sqlStr)
	if err != nil {
		fmt.Println("Error reading table information: ", err.Error())
		return
	}

	defer rows.Close()

	for rows.Next() {
		col := column{}
		err = rows.Scan(&col.ColumnName, &col.Type, &col.Nullable, &col.TableName, &col.ColumnComment)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		//col.Json = strings.ToLower(col.ColumnName)
		col.Json = col.ColumnName
		col.ColumnName = t.camelCase(col.ColumnName)
		col.Type = typeForMysqlToGo[col.Type]

		if t.enableJsonTag {
			col.Json = fmt.Sprintf("`json:\"" + col.Json + "\"`")
		}
		if _, ok := tableColumns[col.TableName]; !ok {
			tableColumns[col.TableName] = []column{}
		}
		tableColumns[col.TableName] = append(tableColumns[col.TableName], col)
	}
	return
}

func (t *Table2Struct) camelCase(str string) string {
	// 是否有表前缀, 设置了就先去除表前缀
	if t.prefix != "" {
		str = strings.Replace(str, t.prefix, "", 1)
	}
	var text string
	for _, p := range strings.Split(str, "_") {
		// 字段首字母大写的同时, 是否要把其他字母转换为小写
		switch len(p) {
		case 0:
		case 1:
			text += strings.ToUpper(p[0:1])
		default:
			// 字符长度大于1时
			//text += strings.ToUpper(p[0:1]) + strings.ToLower(p[1:])
			text += strings.ToUpper(p[0:1]) + p[1:]
		}
	}
	return text
}

// grom单独处理
func FilterToGorm(col column) column {
	if col.ColumnName == "Id" ||
		col.ColumnName == "UserId" {
		//col.ColumnName = "ID"
		col.Type = "uint32"
	}
	if col.ColumnName == "UpdateAt" ||
		col.ColumnName == "UpdatedAt" ||
		col.ColumnName == "ModifyTime" ||
		col.ColumnName == "CreatedAt" ||
		col.ColumnName == "CreateAt" ||
		col.ColumnName == "InsertTime" {
		col.Type = "int64"
	}
	return col
}

/**
 * 获取db名称
 */
func (t *Table2Struct) GetDbName() string {
	dsn := t.dsn
	start := strings.Index(dsn, "/")
	end := strings.Index(dsn, "?")
	return dsn[start+1 : end]
}

/**
 * 表名转换 (例如: user->User userOrder->UserOrder user_order->UserOrder)
 */
func tableToCap(tab string) {

}

/**
 * 生成方法
 */
func (t *Table2Struct) makeOrmMethod() string {
	methodStr := "\r\n"
	// list create update delete one
	methodStr += t.makeOneFuncStr()
	methodStr += t.makeListFuncStr()
	methodStr += t.makeCreateFuncStr()
	methodStr += t.makeUpdateFuncStr()
	methodStr += t.makeDeleteFuncStr()
	return methodStr
}

// todo demo
//func (t *Table2Struct) makeOneFuncStr() string {
//	tableName := t.camelCase(t.table)
//	oneFuncStr := "\r\n"
//	oneFuncStr += "func " + tableName + "One(condition interface{}) (*" + tableName + ",error){\r\n"
//	oneFuncStr += tableName + ":= new(" + tableName + ")\r\n"
//	oneFuncStr += "res, err := " + t.table + ".Where(condition).One()\r\n"
//	oneFuncStr += "if err != nil {\r\n"
//	oneFuncStr += "		if err == sql.ErrNoRows { \r\n return " + tableName + ", nil \r\n}\r\n"
//	oneFuncStr += "return nil, err\r\n"
//	oneFuncStr += "}\r\n"
//	oneFuncStr += "	if err := res.ToStruct(" + tableName + "); err != nil { \r\nreturn nil, err \r\n}\r\n"
//	oneFuncStr += "return " + tableName + ", nil"
//	oneFuncStr += "}\r\n"
//	oneFuncStr += "\r\n"
//	return oneFuncStr
//}

func (t *Table2Struct) makeOneFuncStr() string {
	tableName := t.camelCase(t.table)
	oneFuncStr := "func " + "One" + tableName + "(where map[string]interface{}) (*" + tableName + ",error){\r\n"
	oneFuncStr += "}\r\n"
	oneFuncStr += "\r\n"
	return oneFuncStr
}

func (t *Table2Struct) makeDeleteFuncStr() string {
	tableName := t.camelCase(t.table)
	deleteFuncStr := "func " + "Delete" + tableName + "(where map[string]interface{}) error {\r\n"
	deleteFuncStr += "}\r\n"
	deleteFuncStr += "\r\n"
	return deleteFuncStr
}

func (t *Table2Struct) makeListFuncStr() string {
	tableName := t.camelCase(t.table)
	listFuncStr := "func " + "List" + tableName + "(where map[string]interface{}, page, size int) ([]*" + tableName + ",error){\r\n"
	listFuncStr += "}\r\n"
	listFuncStr += "\r\n"
	return listFuncStr
}

func (t *Table2Struct) makeCreateFuncStr() string {
	tableName := t.camelCase(t.table)
	addFuncStr := "func " + "Create" + tableName + "(where map[string]interface{}) error {\r\n"
	addFuncStr += "}\r\n"
	addFuncStr += "\r\n"
	return addFuncStr
}

func (t *Table2Struct) makeUpdateFuncStr() string {
	tableName := t.camelCase(t.table)
	updateFuncStr := "func " + "Update" + tableName + "(where map[string]interface{}) error {\r\n"
	updateFuncStr += "}\r\n"
	updateFuncStr += "\r\n"
	return updateFuncStr
}
