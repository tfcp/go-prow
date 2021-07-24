package models

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

var (
	ownerTitle    = "Owner列表"
	ownerDescribe = "Owner列表"
	OwnerTable    = "prow_owner"
	owner         = Owner{}
)

type Owner struct {
	Model
	OwnerName   string `json:"owner_name" zh:"owner名称" en:"Owner Name"`
	ProjectName string `json:"project_name" zh:"仓库名称" en:"Project Name"`
	PathName    string `json:"path_name" zh:"路径名称" en:"Path Name"`
	Phone       string `json:"phone" zh:"电话" en:"Phone"`
	Email       string `json:"email" zh:"邮箱" en:"Email"`
	Description string `json:"description" zh:"描述" en:"description"`
}

func GetProwOwnerTable(ctx *context.Context) table.Table {
	prowOwner := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := prowOwner.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int)
	info.AddField(Lan(&owner, "owner_name"), "owner_name", db.Varchar).FieldFilterable()
	info.AddField(Lan(&owner, "project_name"), "project_name", db.Varchar).FieldFilterable()
	info.AddField(Lan(&owner, "path_name"), "path_name", db.Varchar).FieldFilterable()
	info.AddField(Lan(&owner, "phone"), "phone", db.Varchar).FieldFilterable()
	info.AddField(Lan(&owner, "email"), "email", db.Varchar)
	info.AddField(Lan(&owner, "description"), "description", db.Mediumtext)
	info.AddField("创建时间", "create_at", db.Timestamp)
	info.AddField("更新时间", "update_at", db.Timestamp)

	info.SetTable("prow_owner").SetTitle(ownerTitle).SetDescription(ownerDescribe)

	formList := prowOwner.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default)
	formList.AddField(Lan(&owner, "owner_name"), "owner_name", db.Varchar, form.Text)
	formList.AddField(Lan(&owner, "project_name"), "project_name", db.Varchar, form.Text)
	formList.AddField(Lan(&owner, "path_name"), "path_name", db.Varchar, form.Text)
	formList.AddField(Lan(&owner, "phone"), "phone", db.Varchar, form.Text)
	formList.AddField(Lan(&owner, "email"), "email", db.Varchar, form.Email)
	formList.AddField(Lan(&owner, "description"), "description", db.Varchar, form.Text)

	formList.SetTable(OwnerTable).SetTitle(ownerTitle).SetDescription(ownerTitle)

	return prowOwner
}
