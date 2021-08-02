package models

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

var (
	confTitle        = "配置列表"
	confDescribe     = "配置列表"
	confTable        = "prow_conf"
	conf             = Conf{}
	confEnable       = "1"
	confUnable       = "2"
	gitRobotTokenKey = "gitlab-robot-token"
	gitApiKey        = "gitlab-api"
	keyWordsKey      = "key-words"
)

type Conf struct {
	Model
	Key      string `json:"key" zh:"配置名称" en:"ConfKey"`
	Value    string `json:"value" zh:"配置值" en:"ConfValue"`
	Status   string `json:"status" zh:"配置状态" en:"ConfStatus"`
	Describe string `json:"describe" zh:"描述" en:"Describe"`
}

type GitConf struct {
	Token    string `json:"token"`
	ApiAddr  string `json:"api_addr"`
	KeyWords string `json:"key_words"`
}

func GetGitConf() (GitConf, error) {
	var (
		conf    Conf
		gitConf GitConf
	)
	whereToken := map[string]interface{}{
		"status": confEnable,
		"key":    gitRobotTokenKey,
	}
	err := gormDb.Where(whereToken).Find(&conf).Limit(1).Error
	if err != nil {
		return gitConf, err
	}
	gitConf.Token = conf.Value
	whereApi := map[string]interface{}{
		"status": confEnable,
		"key":    gitApiKey,
	}
	err = gormDb.Where(whereApi).Find(&conf).Limit(1).Error
	if err != nil {
		return gitConf, err
	}
	gitConf.ApiAddr = conf.Value
	whereKey := map[string]interface{}{
		"status": confEnable,
		"key":    keyWordsKey,
	}
	err = gormDb.Where(whereKey).Find(&conf).Limit(1).Error
	if err != nil {
		return gitConf, err
	}
	gitConf.KeyWords = conf.Value
	return gitConf, nil
}

func GetProwConfTable(ctx *context.Context) table.Table {

	prowConf := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := prowConf.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField(Lan(&conf, "key"), "key", db.Varchar)
	info.AddField(Lan(&conf, "value"), "value", db.Varchar)
	info.AddField(Lan(&conf, "describe"), "describe", db.Varchar)
	info.AddField(Lan(&conf, "status"), "status", db.Int).FieldBool(confEnable, confUnable)

	info.SetTable(confTable).SetTitle(confTitle).SetDescription(confDescribe)

	formList := prowConf.GetForm()
	formList.AddField(Lan(&conf, "key"), "key", db.Varchar, form.Text).FieldDisplayButCanNotEditWhenUpdate()
	formList.AddField(Lan(&conf, "value"), "value", db.Varchar, form.Text)
	formList.AddField(Lan(&conf, "describe"), "describe", db.Varchar, form.Text)
	formList.AddField(Lan(&conf, "status"), "status", db.Int, form.SelectSingle).FieldOptions(types.FieldOptions{
		{Text: "启用", Value: confEnable},
		{Text: "禁用", Value: confUnable},
	}).
		// 设置默认值
		FieldDefault(confEnable)

	formList.SetTable(confTable).SetTitle(confTitle).SetDescription(confDescribe)

	return prowConf
}
