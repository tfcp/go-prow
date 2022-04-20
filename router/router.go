package router

import (
	"prow/app/api/ci"
	"prow/app/api/demo"
	_ "prow/docs" // gin-swagger
	"prow/internal/middleware/cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	//_ "prow/internal/rice"
	//"prow/library/utils"
	//rice "github.com/GeertJohan/go.rice"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

var (
	Router *gin.Engine
)

func RegisterRouter() {
	Router = gin.Default()
	// pprof
	pprof.Register(Router)
	Router.Use(cors.Cors())
	//fs := utils.EmbeddingFileSystem(rice.MustFindBox(enum.RicePath).HTTPBox())
	//Router.Use(utils.Serve("", fs))
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	dm := Router.Group("demo")
	//dm.Use(jwt.JWT()) # jwt auth
	dm.GET("/hello-list", demo.HelloListApi)
	dm.GET("/hello-info", demo.HelloInfoApi)
	dm.GET("/user-list", demo.UserListApi)
	dm.GET("/user-detail", demo.UserDetailApi)
	dm.POST("/user-delete", demo.UserDeleteApi)
	dm.POST("/user-change", demo.UserChangeStatusApi)
	dm.POST("/user-save", demo.UserSaveApi)
	us := Router.Group("user")
	us.GET("/login", demo.LoginApi)
	us.GET("/info", demo.InfoApi)
	// ci接口
	ciGroup := Router.Group("git/hook")
	{
		// 合并请求
		ciGroup.POST("/merge", ci.MergeApi)
		// 评论
		ciGroup.POST("/comment", ci.CommentApi)
	}

}
