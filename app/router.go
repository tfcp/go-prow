package app

import (
	"github.com/gin-gonic/gin"
	"go-prow/app/api"
)

var globalRouterPrefix = "/prow"

func InitRouter() *gin.Engine {
	r := gin.New()
	gitlab := r.Group(globalRouterPrefix + "/gitlab")
	{
		// owner列表
		gitlab.POST("/merge", api.MergeApi)
		gitlab.POST("/comment", api.CommentApi)
	}

	owner := r.Group(globalRouterPrefix + "/owner")
	{
		// owner列表
		owner.GET("/", api.MergeApi)
		//owner.GET("/comment", api.CommentApi)
	}
	return r
}
