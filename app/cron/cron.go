package cron

import (
	"prow/app/cron/demo"
	"prow/library/utils"
)

var (
	cronStdoutTitle    = []string{"任务名称", "表达式", "负责人", "创建时间"}
	cronStdoutContents = [][]string{}
)

// (Seconds Minutes Hours Day Month Week)
// @hourly @daily @weekly @monthly @yearly
// @every <duration>: @every 1h30m10s 每隔1小时30分钟10秒
// 2 * * * * * 		: 每分钟第2秒执行
// */5 * * * * *	: 每分钟第2秒执行
// 0 * * * * *		: 每分钟执行
func Cron() {
	// cron任务
	demoStdout := demo.CronDemo()

	cronStdoutContents = append(cronStdoutContents, demoStdout...)

	// stdout cron list
	utils.TableStdout(cronStdoutTitle, cronStdoutContents)
}
