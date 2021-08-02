package service

import (
	"encoding/json"
	"fmt"
	"github.com/xanzy/go-gitlab"
	"go-prow/app/consumer"
	"go-prow/app/models"
	"strings"
)

var (
	mergeOpen    = "opened"
	TargetBranch = "master"
)

func getGitClient() (*gitlab.Client, error) {
	gitConf, err := models.GetGitConf()
	if err != nil {
		// todo log
		return nil, err
	}
	git, err := gitlab.NewClient(gitConf.Token, gitlab.WithBaseURL(gitConf.ApiAddr))
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
	// filter
	if mrState != mergeOpen || targetBranch != TargetBranch {
		return nil
	}
	git, err := getGitClient()
	if err != nil {
		// todo log
		return err
	}
	// get changes list
	mrChange, _, err := git.MergeRequests.GetMergeRequestChanges(projectId, mrId, &gitlab.GetMergeRequestChangesOptions{})
	if err != nil {
		// todo log
		return err
	}
	paths, err := getChangePath(mrChange)
	if err != nil {
		// todo log
		return err
	}
	if len(paths) > 0 {
		// notice send
		consumer.SetMergeNotice(paths)
	}
	return err
}

func CommentService(mrState, note, userName, targetBranch string, mrId, proId int) error {
	if mrState != mergeOpen || targetBranch != TargetBranch {
		return nil
	}
	gitConf, err := models.GetGitConf()
	if err != nil {
		// todo log
		return err
	}
	// must keywords
	if strings.ToLower(note) != gitConf.KeyWords {
		return nil
	}
	_, err = getGitClient()
	if err != nil {
		// todo log
		return err
	}
	return nil
}

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

func getChangePath(mr *gitlab.MergeRequest) (changePaths []string, err error) {
	allChange := make([]changeContent, 0)
	if mr.Changes == nil {
		// todo log
		return
	}
	mcs, err := json.Marshal(mr.Changes)
	if err != nil {
		// todo log
		return
	}
	err = json.Unmarshal(mcs, &allChange)
	if err != nil {
		// todo log
		return
	}
	for _, v := range allChange {
		changePaths = append(changePaths, v.OldPath)
	}
	changePaths = ownersPathFilter(changePaths)
	return
}

func ownersPathFilter(changePaths []string) (ownerPaths []string) {
	if len(changePaths) == 0 {
		return
	}
	owners, err := models.GetOwners()
	if err != nil {
		// todo log
		return
	}
	resChangePathMap := map[string]string{}
	for _, owner := range owners {
		for _, changePath := range changePaths {
			if strings.Contains(changePath, owner.PathName) {
				// filter duplicate path
				resChangePathMap[owner.PathName] = owner.PathName
			}
		}
	}
	// get path data
	for _, v := range resChangePathMap {
		ownerPaths = append(ownerPaths, v)
	}
	return
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
