package demo

import (
	"prow/internal/service/demo"
	"prow/library/code"
	"prow/library/utils"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
)

type RequestUserList struct {
	Name     string `json:"name" form:"name"`
	Sex      int    `json:"sex" form:"sex"`
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"page_size" form:"page_size"`
}

// @Summary swagger文档示例2
// @Tags 示例
// @Param name query string true "name " required
// @Param sex query string true "sex " required
// @Param page query string true "page " required
// @Param page_size query string true "page_size " required
// @Success 200 {object} utils.Res
// @Router /demo/user-list [get]
func UserListApi(c *gin.Context) {
	var reqUserList RequestUserList
	c.Bind(&reqUserList)
	if err := gvalid.CheckStruct(c, reqUserList, nil); err != nil {
		utils.Response(c, code.ErrSystem, err.FirstString())
		return
	}
	hs := demo.NewUserService()
	whereCondition := map[string]interface{}{}
	if reqUserList.Name != "" {
		whereCondition["name"] = reqUserList.Name
	}
	if reqUserList.Sex != 0 {
		whereCondition["sex"] = reqUserList.Sex
	}
	ListInfo, err := hs.List(whereCondition, reqUserList.Page, reqUserList.PageSize)
	if err != nil {
		utils.Response(c, code.ErrSystem, err.Error())
		return
	}
	res := map[string]interface{}{
		"result": ListInfo,
	}
	utils.Response(c, code.ErrSuccess, res)
}

type RequestUserDetail struct {
	Id int `json:"id" form:"id" valid:"id      @required#id不能为空"`
}

func UserDetailApi(c *gin.Context) {
	var reqUserDetail RequestUserDetail
	c.Bind(&reqUserDetail)
	if err := gvalid.CheckStruct(c, reqUserDetail, nil); err != nil {
		utils.Response(c, code.ErrSystem, err.FirstString())
		return
	}
	hs := demo.NewUserService()
	whereCondition := map[string]interface{}{}
	whereCondition["id"] = reqUserDetail.Id
	userDetail, err := hs.One(whereCondition)
	if err != nil {
		utils.Response(c, code.ErrSystem, err.Error())
		return
	}
	res := map[string]interface{}{
		"result": userDetail,
	}
	utils.Response(c, code.ErrSuccess, res)
}

type RequestUserDelete struct {
	Id int `json:"id" form:"id" valid:"id      @required#id不能为空"`
}

func UserDeleteApi(c *gin.Context) {
	var reqUserDelete RequestUserDelete
	c.Bind(&reqUserDelete)
	if err := gvalid.CheckStruct(c, reqUserDelete, nil); err != nil {
		utils.Response(c, code.ErrSystem, err.FirstString())
		return
	}
	hs := demo.NewUserService()
	err := hs.Delete(reqUserDelete.Id)
	if err != nil {
		utils.Response(c, code.ErrSystem, err.Error())
		return
	}
	utils.Response(c, code.ErrSuccess, map[string]interface{}{})
}

type RequestUserChange struct {
	Id     int `json:"id" form:"id" valid:"id      @required|integer|min:1#id不能为空"`
	Status int `json:"status" form:"status" valid:"status      @required|integer|min:1#status不能为空"`
}

func UserChangeStatusApi(c *gin.Context) {
	var reqUserChange RequestUserChange
	c.Bind(&reqUserChange)
	if err := gvalid.CheckStruct(c, reqUserChange, nil); err != nil {
		utils.Response(c, code.ErrSystem, err.FirstString())
		return
	}
	hs := demo.NewUserService()
	err := hs.ChangeStatus(reqUserChange.Id, reqUserChange.Status)
	if err != nil {
		utils.Response(c, code.ErrSystem, err.Error())
		return
	}
	utils.Response(c, code.ErrSuccess, map[string]interface{}{})
}

type RequestUserSave struct {
	Id           int    `json:"id" form:"id" valid:"id      @required|integer|min:1#id不能为空"`
	Status       int    `json:"status" form:"status" valid:"status      @required|integer|min:1#status不能为空"`
	Role         int    `json:"role" form:"role" valid:"role      @required|integer|min:1#role不能为空"`
	Sex          int    `json:"sex" form:"sex" valid:"sex      @required|integer|min:1#sex不能为空"`
	Age          int    `json:"age" form:"age"`
	Name         string `json:"name" form:"name" valid:"name      @required#name不能为空"`
	Avatar       string `json:"avatar" form:"avatar"`
	Introduction string `json:"introduction" form:"introduction" valid:"introduction      @required#introduction不能为空"`
}

func UserSaveApi(c *gin.Context) {
	var reqUserSave RequestUserSave
	c.Bind(&reqUserSave)
	if err := gvalid.CheckStruct(c, reqUserSave, nil); err != nil {
		utils.Response(c, code.ErrSystem, err.FirstString())
		return
	}
	hs := demo.NewUserService()
	err := hs.Save(reqUserSave.Id,
		reqUserSave.Status,
		reqUserSave.Role,
		reqUserSave.Sex,
		reqUserSave.Age,
		reqUserSave.Name,
		reqUserSave.Avatar,
		reqUserSave.Introduction)
	if err != nil {
		utils.Response(c, code.ErrSystem, err.Error())
		return
	}
	utils.Response(c, code.ErrSuccess, map[string]interface{}{})
}
