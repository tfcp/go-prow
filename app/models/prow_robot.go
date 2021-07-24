package models

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

var (
	robotTitle    = "机器人列表"
	robotDescribe = "机器人列表"
	RobotTable    = "prow_robot"
	robot         = Robot{}
)

type Robot struct {
	Model
	RobotName   string `json:"robot_name" zh:"仓库名称" en:"Robot Name"`
	GitlabUrl   string `json:"gitlab_url" zh:"gitlab地址" en:"Gitlab Url"`
	GitlabToken string `json:"gitlab_token" zh:"Gitlab Token" en:"Gitlab Token"`
	WatchWord   string `json:"watch_word" zh:"监听内容" en:"Watch Word"`
	Status      int    `json:"status" zh:"状态" en:"Status"`
	Description string `json:"description" zh:"描述" en:"Description"`
}

func GetGitRobot(gitWebUrl, gitSshUrl string) (Robot, error) {
	var robot Robot
	//err := db.Find(&robot, condition).Error
	where := map[string]interface{}{
		"status":     1,
		"gitlab_url": gitWebUrl,
	}
	whereOr := map[string]interface{}{
		"status":     1,
		"gitlab_url": gitSshUrl,
	}
	err := gormDb.Where(where).Or(whereOr).Find(&robot).Limit(1).Error
	if err != nil {
		return robot, err
	}
	return robot, nil
}

func GetProwRobotTable(ctx *context.Context) table.Table {

	prowRobot := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := prowRobot.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField(Lan(&robot, "robot_name"), "robot_name", db.Varchar)
	info.AddField(Lan(&robot, "description"), "description", db.Mediumtext)
	info.AddField(Lan(&robot, "gitlab_url"), "gitlab_url", db.Varchar)
	info.AddField(Lan(&robot, "gitlab_token"), "gitlab_token", db.Varchar)
	info.AddField(Lan(&robot, "status"), "status", db.Int)
	info.AddField("创建时间", "create_at", db.Timestamp)
	info.AddField("更新时间", "update_at", db.Timestamp)

	info.SetTable(RobotTable).SetTitle(robotTitle).SetDescription(robotDescribe)

	formList := prowRobot.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default)
	formList.AddField(Lan(&robot, "robot_name"), "robot_name", db.Varchar, form.Text)
	//formList.AddField(Lan(&robot, "description"), "description", db.Mediumtext, form.RichText)
	formList.AddField(Lan(&robot, "description"), "description", db.Varchar, form.Text)
	formList.AddField(Lan(&robot, "gitlab_token"), "gitlab_token", db.Varchar, form.Text)
	formList.AddField(Lan(&robot, "status"), "status", db.Int, form.Text)

	formList.SetTable(RobotTable).SetTitle(robotTitle).SetDescription(robotDescribe)

	return prowRobot
}
