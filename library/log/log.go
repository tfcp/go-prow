package log

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
)

var Logger *glog.Logger

func Setup() {
	Logger = glog.New()
	if err := Logger.SetConfigWithMap(g.Map{
		"path":   g.Config().GetString("log.path"),
		"level":  g.Config().GetString("log.level"),
		"stdout": g.Config().GetBool("log.stdout"),
	}); err != nil {
		Logger.Fatal(err)
	}
}
