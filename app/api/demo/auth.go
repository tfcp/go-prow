package demo

import (
	"prow/internal/service/demo"
	"prow/library/code"
	"prow/library/utils"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
)

type RequestLogin struct {
	Name string `json:"username" form:"username" valid:"username      @required#username不能为空"`
	Pwd  string `json:"password" form:"password" valid:"password      @required#password不能为空"`
}

func LoginApi(c *gin.Context) {
	var reqLogin RequestLogin
	c.Bind(&reqLogin)
	if err := gvalid.CheckStruct(c, reqLogin, nil); err != nil {
		utils.Response(c, code.ErrSystem, err.FirstString())
		return
	}
	us := demo.UserService{}
	whereMap := map[string]interface{}{
		"name": reqLogin.Name,
		"pwd":  reqLogin.Pwd,
	}
	um, err := us.One(whereMap)
	if err != nil {
		utils.Response(c, code.ErrPwd, err.Error())
		return
	}
	token, err := utils.GenerateToken(um.Name, um.Avatar, um.Introduction, um.Role)
	if err != nil {
		utils.Response(c, code.ErrPwd, err.Error())
		return
	}
	utils.Response(c, code.ErrSuccess, map[string]interface{}{
		"token": token,
	})
}

type RequestInfo struct {
	Token string `json:"token" form:"token" valid:"token      @required#token不能为空"`
}

func InfoApi(c *gin.Context) {
	var reqInfo RequestInfo
	c.Bind(&reqInfo)
	if err := gvalid.CheckStruct(c, reqInfo, nil); err != nil {
		utils.Response(c, code.ErrSystem, err.FirstString())
		return
	}
	//token := c.Query("token")
	hs := demo.NewUserService()
	oneInfo, err := hs.Info(reqInfo.Token)
	//oneInfo, err := hs.Info(token)
	if err != nil {
		utils.Response(c, code.ErrSystem, err.Error())
		return
	}
	utils.Response(c, code.ErrSuccess, oneInfo)
}

func LogoutApi() {

}
