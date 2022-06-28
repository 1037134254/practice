package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Ping
// @Tags 连接测试
// @Summary 连接
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msessage": "Pong"})
}
