package service

import (
	"fmt"
	"github.com/dxvgef/limiter"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

//func NewSpeedLimiter(r float64, b int) *rate.Limiter {
//	return rate.NewLimiter(rate.Limit(r), int(b))
//}

// RatePut
// @Tags 速度改变
// @Summary 速度
// @Param speed query int false "speed"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /problem-list [get]

//func RatePut(c *gin.Context) {
//	speeds := c.Query("speed")
//	c.JSON(http.StatusOK, map[string]interface{}{
//		"msg":   "修改速度成功",
//		"speed": speed,
//	})
//}

func Demo(c *gin.Context) {
	// 需要搞一个副本
	copyContext := c.Copy()
	c.JSON(http.StatusOK, "hehe")
	// 异步处理
	go func() {
		time.Sleep(5 * time.Second)
		log.Println("异步执行：" + copyContext.Request.URL.Path)
	}()
}

var speed string

// Rate
// @Tags 文件
// @Summary 文件下载速率控制
// @Param path query string true "path"
// @Param fileName query string true "fileName"
// @Param speed query string true "speed"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /rate [post]
func Rate(c *gin.Context) {
	resp := c.Writer
	req := c.Request
	//path := c.Query("path")
	//fileName := c.Query("fileName")
	//speed := c.Query("speed")
	path := c.PostForm("path")
	fileName := c.PostForm("fileName")
	speed = c.PostForm("speed")
	float, err := strconv.ParseFloat(speed, 10)
	fmt.Printf("当前下载速度为:%v\n", float)
	//if float < 100 {
	//	log.Printf("限速过低")
	//	return
	//}
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; fileName=%s", fileName))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	if err != nil {
		c.JSON(500, gin.H{"msg": "Error ParseFloat:" + err.Error()})
		return
	}
	//rate.NewLimiter(10,100)
	// speed 控制下载速度 最少100KB ->100
	// l 控制多少时间内放入的令牌数量
	// b 最终能放多少令牌数量
	if err := limiter.ServeFile(resp, req, path+"\\"+fileName, float*1024); err != nil {
		resp.WriteHeader(500)
		resp.Write([]byte(err.Error()))
	}
	log.Printf("发送成功")
}

func updateSpeed(c *gin.Context) {
	query := c.Query("speed")
	if query != speed {

	} else {
		c.JSON(http.StatusOK, gin.H{"msg": "速度一样不修改"})
		return
	}
}
