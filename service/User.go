package service

import (
	"example.com/m/v2/defaults"
	"example.com/m/v2/helper"
	"example.com/m/v2/models"
	"github.com/gin-gonic/gin"
	"gorm/gorm"
	"log"
	"net/http"
	"strconv"
	"time"
)

// GetUserDetail
// @Tags 用户
// @Summary 用户详情
// @Param identity query string false "problem_identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /user-detail [get]
func GetUserDetail(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户唯一标识符不能为空",
		})
		return
	}
	date := new(models.UserBasic)
	err := models.DB.Where("?", identity).Find(&date).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "预料之外的错误" + identity + "Error" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  date,
	})
	return
}

// SendCode
// @Tags 用户
// @Summary 发送验证码
// @Param email formData string false "email"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /send-code [post]
func SendCode(c *gin.Context) {
	email := c.PostForm("email")
	if email == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "邮箱为空",
		})
		return
	}
	code := helper.GetRand()
	err := helper.SendCode(email, code)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "验证码不正确重新获取验证码",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg": gin.H{
			"验证码":   code,
			"email": email,
		},
	})
	// 验证码设入redis
	models.RDB.Set(c, email, code, time.Second*300)
	return
}

// Login
// @Tags 用户
// @Summary 用户登录
// @Param username formData string false "username"
// @Param password formData string false "password"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /login [post]
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户名或密码为空",
		})
		return
	}
	password = helper.MD5(password)
	date := new(models.UserBasic)
	err := models.DB.Where("name = ? and password = ?", username, password).First(&date).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户名或密码错误",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get UserBasic Error:" + err.Error(),
		})
		return
	}
	token, err := helper.GenerateToken(date.Identity, date.Name, date.IsAdmin)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "GenerateToken Error:" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"date": map[string]interface{}{
			"token": token,
		},
	})
}

// Register
// @Tags 用户
// @Summary 用户注册
// @Param email formData string true "email"
// @Param code formData string true "code"
// @Param name formData string true "name"
// @Param password formData string true "password"
// @Param phone formData string false "phone"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /register [post]
func Register(c *gin.Context) {
	email := c.PostForm("email")
	userCode := c.PostForm("code")
	name := c.PostForm("name")
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	if email == "" || userCode == "" || name == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不正确",
		})
		return
	}
	sysCode, err := models.RDB.Get(c, email).Result()
	if err != nil {
		log.Printf("Get Code Error:%v\n", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "验证码不正确重新获取",
		})
		return
	}
	if sysCode != userCode {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "验证码不正确",
		})
		return
	}

	var cnt int64
	err = models.DB.Where("mail=?", email).Model(new(models.UserBasic)).Count(&cnt).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get User1 Error:" + err.Error(),
		})
		return
	}
	if cnt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "该邮箱已注册",
		})
		return
	}
	Identity := helper.UUid()
	// 插入数据
	date := &models.UserBasic{
		Identity: Identity,
		Name:     name,
		Password: helper.MD5(password),
		Phone:    phone,
		Mail:     email,
	}
	err = models.DB.Create(date).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Create User2 Error" + err.Error(),
		})
		return
	}
	// 生成 token
	token, err := helper.GenerateToken(Identity, name, date.IsAdmin)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Create Token Error" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"date": map[string]interface{}{"token": token},
	})
	return
}

// GetRankList
// @Tags 用户
// @Summary 用户排行
// @Param page query int false "page"
// @Param size query int false "size"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /rank-list [post]
func GetRankList(c *gin.Context) {
	size, _ := strconv.Atoi(c.DefaultQuery("size", defaults.DefaultSize))
	page, err := strconv.Atoi(c.DefaultQuery("size", defaults.DefaultPage))
	if err != nil {
		log.Printf("GetRankList:" + err.Error())
	}
	page = (page - 1) * size

	var count int64
	list := make([]*models.UserBasic, 0)
	err = models.DB.Model(new(models.UserBasic)).Count(&count).Order("finis_problem_num desc,submit_num asc").
		Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get Rank List Error" + err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"date": map[string]interface{}{
			"list":  list,
			"count": count,
		},
	})
}
