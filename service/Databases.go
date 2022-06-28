package service

import (
	"example.com/m/v2/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm/mysql"
)

// Databases
// @Tags 数据源
// @Summary 数据库初始化
// @Param email formData string false "user"
// @Param email formData string false "password"
// @Param email formData string false "ipAndPort"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /data [post]
func Databases(c *gin.Context) {
	user := c.PostForm("user")
	password := c.PostForm("password")
	ipAndPort := c.PostForm("ipAndPort")
	open, err := gorm.Open(mysql.Open(user+":"+password+"@("+ipAndPort+")/OnlinePractice?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		return
	}
	open.AutoMigrate(models.ProblemBasic{})
	open.AutoMigrate(models.ProblemCategory{})
	open.AutoMigrate(models.SubmitBasic{})
	open.AutoMigrate(models.UserBasic{})
	err = open.AutoMigrate(models.TestCase{})

}
