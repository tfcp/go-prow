package gitlab

import (
	"fmt"
	"prow/library/log"
	"github.com/xanzy/go-gitlab"
)

func GetGitLabClient() (*gitlab.Client, error) {
	git, err := gitlab.NewClient(setting.GitLabSetting.Token, gitlab.WithBaseURL(setting.GitLabSetting.Url))
	if err != nil {
		log.Logger.Error("CommentService error:", err)
		return nil, err
	}
	return git, nil
}

func getLockKey(proId, mrId int) string {
	return fmt.Sprintf("git_link_%v_%v", proId, mrId)
}

func getTmplKey(proId, mrId int) string {
	return fmt.Sprintf("git_tmpl_%v_%v", proId, mrId)
}
