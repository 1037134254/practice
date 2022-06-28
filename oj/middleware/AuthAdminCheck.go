package middleware

import (
	"example.com/m/v2/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthAdminCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		// 解析token
		token, err := helper.AnalyseToken(auth)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "AuthAdminCheck Authorization ",
			})
			return
		}
		if token.IsAdmin != 1 || token == nil {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "AuthAdminCheck Unauthorized2",
			})
			return
		}
		c.Next()
	}
}
