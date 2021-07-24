package models

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

var (
	pathTitle    = "路径列表"
	pathDescribe = "路径列表"
	PathTable    = "prow_path"
	path         = Path{}
)

type Path struct {
	Model
	PathName    string `json:"path_name" zh:"路径名称" en:"Path Name"`
	ProjectName string `json:"project_name" zh:"仓库名称" en:"Project Name"`
	Description string `json:"description" zh:"描述" en:"description"`
}

func GetProwPathTable(ctx *context.Context) table.Table {

	prowPath := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := prowPath.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField(Lan(&path, "path_name"), "path_name", db.Varchar)
	info.AddField(Lan(&path, "project_name"), "project_name", db.Varchar)
	info.AddField(Lan(&path, "description"), "description", db.Mediumtext)
	info.AddField("Create_at", "create_at", db.Timestamp)
	info.AddField("Update_at", "update_at", db.Timestamp)

	info.SetTable(PathTable).SetTitle(pathTitle).SetDescription(pathDescribe)

	formList := prowPath.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default)
	formList.AddField(Lan(&path, "path_name"), "path_name", db.Varchar, form.Text)
	formList.AddField(Lan(&path, "project_name"), "project_name", db.Varchar, form.Text)
	formList.AddField(Lan(&path, "description"), "description", db.Varchar, form.Text)

	formList.SetTable(PathTable).SetTitle(pathTitle).SetDescription(pathDescribe)

	return prowPath
}
