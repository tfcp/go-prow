package models

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

var (
	projectTitle    = "仓库列表"
	projectDescribe = "仓库列表"
	ProjectTable    = "prow_project"
	project         = Project{}
)

type Project struct {
	Model
	ProjectName string `json:"project_name" zh:"仓库名称" en:"Project Name"`
	Description string `json:"description" zh:"描述" en:"Description"`
	Url         string `json:"url" zh:"仓库地址" en:"url"`
}

func GetProwProjectTable(ctx *context.Context) table.Table {

	prowProject := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := prowProject.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField(Lan(&project, "project_name"), "project_name", db.Varchar)
	info.AddField(Lan(&project, "url"), "url", db.Mediumtext)
	info.AddField(Lan(&project, "description"), "description", db.Mediumtext)
	info.AddField("创建时间", "create_at", db.Timestamp)
	info.AddField("更新时间", "update_at", db.Timestamp)

	info.SetTable(ProjectTable).SetTitle(projectTitle).SetDescription(projectDescribe)

	formList := prowProject.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default)
	formList.AddField(Lan(&project, "project_name"), "project_name", db.Varchar, form.Text)
	formList.AddField(Lan(&project, "description"), "description", db.Varchar, form.Text)

	formList.SetTable(ProjectTable).SetTitle(projectTitle).SetDescription(projectDescribe)

	return prowProject
}
