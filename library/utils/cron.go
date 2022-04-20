package utils

import (
	"prow/library/log"
	"github.com/gogf/gf/os/gcron"
	"time"
)

func AddCron(expression, cronName, owner string, cronFunc func()) []string {
	_, err := gcron.Add(expression, func() {
		cronFunc()
	}, cronName)
	if err != nil {
		log.Logger.Errorf("%s CronStart Error: %v", cronName, err)
		return []string{}
	}
	return []string{cronName, expression, owner, time.Now().Format("2006-01-02 15:04:05")}
}
