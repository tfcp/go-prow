package ci

import (
	"encoding/json"
	"github.com/xanzy/go-gitlab"
	"prow/internal/enum"
	"prow/internal/model/ci"
	gitlab2 "prow/library/gitlab"
	"prow/library/log"
	"strings"
)

type GitService struct {
	BaseCi
}

func NewGitService() (s *GitService) {
	s = &GitService{}
	s.ownerModel = &ci.Owner{}
	return s
}

func (this *GitService) Comment(mrState, content, userName, targetBranch string, mrId, proId int) error{
	var err error
	// merge状态必须为提交状态并且是合并master才触发
	if mrState != enum.MergeOpen || targetBranch != enum.TargetBranch {
		return err
	}

	// 必须为准许内容
	if strings.ToLower(content) != enum.Approve {
		return err
	}

	git, err := gitlab2.GetGitLabClient()
	if err != nil {
		log.Logger.Error("CommentService error:", err)
		return err
	}
	m, _, err := git.MergeRequests.GetMergeRequestChanges(proId, mrId, &gitlab.GetMergeRequestChangesOptions{})
	if err != nil {
		log.Logger.Error("CommentService error:", err)
		return err
	}
	approve, paths, ownerPaths := this.isOwnerApprove(m, userName)
	// review机器人回复内容
	if err := Reply(m.ProjectID, m.IID, paths, ownerPaths); err != nil {
		log.Logger.Error("reply error:", err)
	}
	//if approve && !isDenyMergeLock(m.ProjectID, m.IID) {
	//	//满足条件 合并代码
	//	_, _, err = git.MergeRequests.AcceptMergeRequest(proId, mrId, &gitlab.AcceptMergeRequestOptions{})
	//	if err != nil {
	//		log.Logger.Error("CommentService AcceptMergeRequestError:", err)
	//		return err
	//	}
	//	// 只通知一次 合并人为第一个reviewer
	//	title := "【大仓代码合并成功】"
	//	if err := DingDing(m.WebURL, userName, "med-common", title, paths, mergedAlertTmpl); err != nil {
	//		log.Logger.Error("DingDing error:", err)
	//	}
	//	// 合并机器人回复内容(回复已合并)
	//	if err := Reply(m.ProjectID, m.IID, []string{}, []string{}); err != nil {
	//		log.Logger.Error("reply error:", err)
	//	}
	//	// 清除回复模板
	//	if _, err := DeleteTmpl(m.ProjectID, m.IID); err != nil {
	//		log.Logger.Error("DeleteTmpl error:", err)
	//	}
	//}
	return err
}

func (this *GitService) Merge(projectId, mrId int, mrState, targetBranch string) error{
	var err error
	// merge状态必须为提交状态并且目标分支为master
	if mrState != enum.MergeOpen || targetBranch != enum.TargetBranch {
		return err
	}
	git, err := gitlab2.GetGitLabClient()
	if err != nil {
		log.Logger.Error("MergeService NewClientError:", err)
		return err
	}
	m, _, err := git.MergeRequests.GetMergeRequestChanges(projectId, mrId, &gitlab.GetMergeRequestChangesOptions{})
	if err != nil {
		log.Logger.Error("MergeService GetMergeRequestChangesError:", err)
		return err
	}
	// 获取目标项目列表
	paths, err := this.getPaths(m)
	if err != nil {
		return err
	}
	// 如果没有属于prow权限范围内的项目就不做提示
	if len(paths) > 0 {
		// todo 通知
	}
	return err
}

func (this *GitService) isOwnerApprove(mr *gitlab.MergeRequest, userName string) (res bool, paths, ownerPath []string) {
	all := make([]ci.ChangeContent, 0)
	if mr.Changes == nil {
		log.Logger.Error("MergeService error: changes is nil")
		return
	}

	mcs, err := json.Marshal(mr.Changes)
	if err != nil {
		log.Logger.Error("MergeService Marshal error:", err)
		return
	}
	err = json.Unmarshal(mcs, &all)
	if err != nil {
		log.Logger.Error("MergeService Unmarshal error:", err)
		return
	}
	changePaths := make([]string, 0)
	for _, v := range all {
		changePaths = append(changePaths, v.OldPath)
	}
	ownerRes, err := this.ownerModel.GetOwners(nil)
	if err != nil {
		log.Logger.Error("isOwnerApprove GetOwners error:", err)
		return
	}
	paths, err = this.getPaths(mr)
	if err != nil {
		return
	}
	// 给项目列表加锁
	this.lockPaths(mr.ProjectID, mr.IID, paths)
	ownerPath = make([]string, 0)
	for _, owner := range ownerRes {
		for _, path := range paths {
			// 1. 确认提交者为owner身份 2.该owner确实有mr项目的权限
			if owner.Name == userName && owner.Path == path {
				res = true
				// 如果这个owner拥有本次提交修改的多个项目的owner 那么 他只需要审核一次即可
				ownerPath = append(ownerPath, owner.Path)
			}
		}
	}
	if res == true {
		//merge解锁
		for _, targetPath := range ownerPath {
			// 解锁相关merge
			this.unlockPath(mr.ProjectID, mr.IID, targetPath)
		}
	}
	return
}
