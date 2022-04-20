package main

import (
	"prow/bootstrap"
	"prow/library/log"
	"github.com/gogf/gf/os/gcmd"
)

func main() {
	if err := gcmd.BindHandleMap(map[string]func(){
		"server":  bootstrap.Run,
		"process": bootstrap.RunProcess,
		"cron":    bootstrap.RunCron,
		"help":    bootstrap.RunHelp,
		"tools":   bootstrap.RunTools,
	}); err != nil {
		log.Logger.Fatal(err)
	}
	if err := gcmd.AutoRun(); err != nil {
		log.Logger.Fatal(err)
	}
}
