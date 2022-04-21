package ci

import (
	"encoding/json"
	"github.com/xanzy/go-gitlab"
	"prow/internal/model/ci"
	gitlab2 "prow/library/gitlab"
	"prow/library/gredis"
	"prow/library/log"
	"strings"
)

type BaseCi struct {
	ownerModel *ci.Owner
}

func (this *BaseCi) getPaths(m *gitlab.MergeRequest) (paths []string, err error) {
	all := make([]ci.ChangeContent, 0)
	if m.Changes == nil {
		log.Logger.Error("MergeService error: changes is nil")
		return
	}

	mcs, err := json.Marshal(m.Changes)
	if err != nil {
		log.Logger.Error("MergeService Marshal error:", err)
		return
	}
	err = json.Unmarshal(mcs, &all)
	if err != nil {
		log.Logger.Error("MergeService Unmarshal error:", err)
		return
	}
	pathRes := make([]string, 0)
	for _, v := range all {
		pathRes = append(pathRes, v.OldPath)
	}
	paths = this.pathsFilter(pathRes)
	return
}

// 过滤项目路径
func (this *BaseCi) pathsFilter(paths []string) (pathsRes []string) {
	if len(paths) == 0 {
		return
	}
	prosRes, err := this.ownerModel.GetOwners(nil)
	if err != nil {
		log.Logger.Error("MergeService pathsFilter error:", err)
		return
	}
	// 过滤数据
	resMap := map[string]string{}
	for _, v1 := range prosRes {
		for _, v2 := range paths {
			if strings.Contains(v2, v1.Path) {
				// 兼容library (有些项目内部也有library 这种需要忽略)
				if v1.Path == "library" && strings.Contains(v2, "/library/") {
					continue
				}
				resMap[v1.Path] = v1.Path
			}
		}
	}
	for _, v := range resMap {
		pathsRes = append(pathsRes, v)
	}
	return
}

// merge项目加锁
func (this *BaseCi) lockPaths(proId, mrId int, paths []string) (res bool) {
	if proId == 0 || mrId == 0 {
		return
	}
	reviewKey := gitlab2.GetLockKey(proId, mrId)
	// 加锁
	if gredis.Exists(reviewKey) {
		return
	}
	for _, v := range paths {
		// 给mr中每个涉及项目加锁
		_, err := gredis.RedisDo("hset", reviewKey, v, 1)
		if err != nil {
			log.Logger.Errorf("redis hset error: %v", err)
			return
		}
	}
	return true
}

// merge项目解锁
func (this *BaseCi) unlockPath(proId, mrId int, path string) (res bool) {
	res = true
	reviewKey := gitlab2.GetLockKey(proId, mrId)
	_, err := gredis.RedisDo("hdel", reviewKey, path)
	if err != nil {
		res = false
	}
	return res
}

// 判断merge是否被锁 (只有涉及的owner全部同意 才可以merge)
func (this *BaseCi) isDenyMergeLock(proId, mrId int) bool {
	return gredis.Exists(gitlab2.GetLockKey(proId, mrId))
}
