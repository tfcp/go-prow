package service

import (
	"encoding/json"
	"fmt"
	"github.com/xanzy/go-gitlab"
	"go-prow/app/consumer"
	"go-prow/app/models"
)

var (
	gitlabApi    = "https://git.medlinker.com/api/v4"
	mergeOpen    = "opened"
	TargetBranch = "master"
)

// gitlab change content
type changeContent struct {
	OldPath     string `json:"old_path"`
	NewPath     string `json:"new_path"`
	AMode       string `json:"a_mode"`
	BMode       string `json:"b_mode"`
	Diff        string `json:"diff"`
	NewFile     bool   `json:"new_file"`
	RenamedFile bool   `json:"renamed_file"`
	DeletedFile bool   `json:"deleted_file"`
}

func getGitClient(gitHttpUrl, gitSshUrl string) (*gitlab.Client, error) {
	robot, err := models.GetGitRobot(gitHttpUrl, gitSshUrl)
	if err != nil {
		// todo log
		return nil, err
	}
	git, err := gitlab.NewClient(robot.GitlabToken, gitlab.WithBaseURL(gitlabApi))
	if err != nil {
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

func MergeService(projectId, mrId int, mrState, targetBranch, gitHttpUrl, gitSshUrl string) error {
	// merge状态必须为提交状态并且目标分支为master
	if mrState != mergeOpen || targetBranch != TargetBranch {
		return nil
	}
	git, err := getGitClient(gitHttpUrl, gitSshUrl)
	//opt := &gitlab.UpdateSettingsOptions{}
	//opt1 := &gitlab.RequestOptionFunc()
	//options ...RequestOptionFunc
	//a, _ := git.Settings.GetSettings()
	//a.PushEventHooksLimit
	if err != nil {
		// todo log
		return err
	}
	// get changes list
	//mr, _, err := git.MergeRequests.GetMergeRequestChanges(projectId, mrId, &gitlab.GetMergeRequestChangesOptions{})
	_, _, err = git.MergeRequests.GetMergeRequestChanges(projectId, mrId, &gitlab.GetMergeRequestChangesOptions{})
	if err != nil {
		// todo log
		return err
	}
	return err
}

func CommentService() {
	consumer.SetNotice(666)
}

func getOwnerChangePath(mr *gitlab.MergeRequest) (paths []string, err error) {
	all := make([]changeContent, 0)
	if mr.Changes == nil {
		//todo log
		//logging.Error("MergeService error: changes is nil")
		return
	}

	mcs, err := json.Marshal(mr.Changes)
	if err != nil {
		//todo log
		//logging.Error("MergeService Marshal error:", err)
		return
	}
	err = json.Unmarshal(mcs, &all)
	if err != nil {
		//todo log
		//logging.Error("MergeService Unmarshal error:", err)
		return
	}
	pathRes := make([]string, 0)
	for _, v := range all {
		pathRes = append(pathRes, v.OldPath)
	}
	//paths = pathsFilter(pathRes)
	return
}
