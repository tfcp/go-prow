package jwt

import (
	code2 "prow/library/code"
	"prow/library/utils"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var msg string
		var data interface{}

		code = code2.ErrSuccess.Code
		msg = code2.ErrSuccess.Message
		token := c.Request.Header.Get("X-Token")
		if token == "" {
			code = code2.ErrParam.Code
			msg = code2.ErrParam.Message
		} else {
			_, err := utils.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = code2.ErrAuthTimeout.Code
					msg = code2.ErrAuthTimeout.Message
				default:
					code = code2.ErrUnauthorized.Code
					msg = code2.ErrUnauthorized.Message
				}
			}
		}

		if code != code2.ErrSuccess.Code {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  msg,
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
