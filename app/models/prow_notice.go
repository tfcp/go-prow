package models

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetProwNoticeTable(ctx *context.Context) table.Table {

	prowNotice := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := prowNotice.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField("Callback_url", "callback_url", db.Varchar)
	info.AddField("Create_at", "create_at", db.Timestamp)
	info.AddField("Update_at", "update_at", db.Timestamp)

	info.SetTable("prow_notice").SetTitle("ProwNotice").SetDescription("ProwNotice")

	formList := prowNotice.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default)
	formList.AddField("Callback_url", "callback_url", db.Varchar, form.Text)

	formList.SetTable("prow_notice").SetTitle("ProwNotice").SetDescription("ProwNotice")

	return prowNotice
}
