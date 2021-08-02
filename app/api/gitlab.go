package api

import (
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/gin-gonic/gin"
	"go-prow/app/service"
	"go-prow/library/util"
	"net/http"
)

var (
	pageSizeList = []string{"10", "50", "100"}
	defaultSize  = 10
)

type ObjectAttributes struct {
	Iid             int    `json:"iid" form:"iid"`
	Note            string `json:"note" form:"note"`
	State           string `json:"state" form:"state"`
	TargetProjectId int    `json:"target_project_id" form:"target_project_id"`
	TargetBranch    string `json:"target_branch" form:"target_branch"`
}

type Project struct {
	Name       string `json:"name" form:"name"`
	GitHttpUrl string `json:"git_http_url" form:"git_http_url"`
	GitSshUrl  string `json:"git_ssh_url" form:"git_ssh_url"`
}

type MergeRequest struct {
	Iid          int    `json:"iid" form:"iid"`
	TargetBranch string `json:"target_branch" form:"target_branch"`
	State        string `json:"state" form:"state"`
}

type User struct {
	Name string `json:"name"`
}

type MergeApiReq struct {
	ObjectAttributes ObjectAttributes `form:"object_attributes" json:"object_attributes"`
	Project          Project          `json:"project" form:"project"`
}

//merge api
func MergeApi(ctx *gin.Context) {
	appG := util.GinProw{C: ctx}
	mq := &MergeApiReq{}
	err := appG.RequestBody(ctx.Request.Body, mq)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}
	service.MergeService(
		mq.ObjectAttributes.TargetProjectId,
		mq.ObjectAttributes.Iid,
		mq.ObjectAttributes.State,
		mq.ObjectAttributes.TargetBranch,
		mq.Project.GitHttpUrl,
		mq.Project.GitSshUrl,
	)
	appG.Response(http.StatusOK, e.SUCCESS, mq)
}

type CommentApiReq struct {
	ProjectId        int              `json:"project_id" form:"project_id"`
	User             User             `json:"user" form:"user"`
	ObjectAttributes ObjectAttributes `form:"object_attributes" json:"object_attributes"`
	MergeRequest     MergeRequest     `form:"merge_request" json:"merge_request"`
}

// comment api
func CommentApi(ctx *gin.Context) {
	appG := util.GinProw{C: ctx}
	cq := &CommentApiReq{}
	err := appG.RequestBody(ctx.Request.Body, cq)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}
	service.CommentService(
		cq.MergeRequest.State,
		cq.ObjectAttributes.Note,
		cq.User.Name,
		cq.MergeRequest.TargetBranch,
		cq.MergeRequest.Iid,
		cq.ProjectId,
	)
	appG.Response(http.StatusOK, e.SUCCESS, cq)
}
