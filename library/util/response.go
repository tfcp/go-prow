package util

import (
	"encoding/json"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/gin-gonic/gin"
	"html/template"
	"io"
	"io/ioutil"
)

type GinProw struct {
	C *gin.Context
}

type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (g *GinProw) RequestBody(bodyReq io.ReadCloser, req interface{}) error {
	body, _ := ioutil.ReadAll(bodyReq)
	if body != nil {
		if err := json.Unmarshal(body, &req); err != nil {
			return err
		}
	}
	return nil
}

// Response setting gin.JSON
func (g *GinProw) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}

// Response goAdmin html panel
func (g *GinProw) ResponsePanel(content, title, description string) (types.Panel, error) {
	return types.Panel{
		// Content 是页面主题内容，为template.html类型
		Content: template.HTML(content),
		// Title 与 Description是标题与描述
		Title:       template.HTML(title),
		Description: template.HTML(description),
	}, nil
}
