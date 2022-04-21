package gitlab

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"prow/library/log"
	"github.com/xanzy/go-gitlab"
)

func GetGitLabClient() (*gitlab.Client, error) {
	git, err := gitlab.NewClient(g.Config().GetString("gitlab.token"), gitlab.WithBaseURL(g.Config().GetString("gitlab.url")))
	if err != nil {
		log.Logger.Error("GetGitLabClient error:", err)
		return nil, err
	}
	return git, nil
}

func GetLockKey(proId, mrId int) string {
	return fmt.Sprintf("git_link_%v_%v", proId, mrId)
}

func GetTmplKey(proId, mrId int) string {
	return fmt.Sprintf("git_tmpl_%v_%v", proId, mrId)
}
