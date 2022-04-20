package utils

import (
	"prow/library/code"
	"github.com/gin-gonic/gin"
)

type Res struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//func Response(Ctx *gin.Context, code int, msg string, data interface{}) {
func Response(Ctx *gin.Context, ErrorMsg *code.Error, data interface{}) {
	Ctx.JSON(200, Res{
		Code:    ErrorMsg.Code,
		Message: ErrorMsg.Message,
		Data:    data,
	})
	return
}
