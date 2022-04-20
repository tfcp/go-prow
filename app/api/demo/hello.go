package demo

import (
	"prow/internal/service/demo"
	"prow/library/code"
	"prow/library/utils"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
)

type RequestHelloInfo struct {
	Name string `json:"name" form:"name" valid:"name      @required#name不能为空"`
}

// @Summary swagger文档示例1
// @Tags 示例
// @Param name query string true "name 名称" required
// @Success 200 {object} utils.Res
// @Router /demo/hello-info [get]
func HelloInfoApi(c *gin.Context) {
	var reqHelloInfo RequestHelloInfo
	c.Bind(&reqHelloInfo)
	if err := gvalid.CheckStruct(c, reqHelloInfo, nil); err != nil {
		utils.Response(c, code.ErrSystem, err.FirstString())
		return
	}
	hs := demo.NewHelloService()
	whereCondition := map[string]interface{}{
		"name": reqHelloInfo.Name,
	}
	oneInfo, err := hs.One(whereCondition)
	if err != nil {
		utils.Response(c, code.ErrSystem, err.Error())
		return
	}
	res := map[string]interface{}{
		"result": oneInfo,
	}

	utils.Response(c, code.ErrSuccess, res)
}

type RequestHelloList struct {
	Name string `json:"name" form:"name" valid:"name      @required#name不能为空"`
}

func HelloListApi(c *gin.Context) {
	var reqHelloList RequestHelloList
	c.Bind(&reqHelloList)
	if err := gvalid.CheckStruct(c, reqHelloList, nil); err != nil {
		utils.Response(c, code.ErrSystem, err.FirstString())
		return
	}
	hs := demo.NewHelloService()
	whereCondition := map[string]interface{}{
		"name": reqHelloList.Name,
	}
	ListInfo, err := hs.List(whereCondition)
	if err != nil {
		utils.Response(c, code.ErrSystem, err.Error())
		return
	}
	res := map[string]interface{}{
		"result": ListInfo,
	}
	utils.Response(c, code.ErrSuccess, res)
}
