package router

import (
	_ "example.com/m/v2/docs"
	"example.com/m/v2/middleware"
	"example.com/m/v2/service"
	"github.com/dxvgef/limiter"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func Router() *gin.Engine {
	engine := gin.Default()
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	engine.GET("/ping", service.Ping)
	// 路由

	// 问题
	engine.GET("/problem-list", service.GetProblemList)
	engine.GET("/problem-detail", service.GetProblemDetail)

	// 用户
	engine.GET("/user-detail", service.GetUserDetail)
	engine.POST("/login", service.Login)
	engine.POST("/send-code", service.SendCode)
	engine.POST("/register", service.Register)
	// 用户排行帮
	engine.POST("/rank-list", service.GetRankList)
	// 提交记录
	engine.GET("/submit-list", service.GetSubmitList)

	// 管理员私有方法
	//engine.POST("/problem-create", service.ProblemCreate)
	check := middleware.AuthAdminCheck()
	group := engine.Group("", check)
	// 分类列表
	group.GET("/category-list", service.GetCategoryList)
	// 分类新增
	group.POST("/category-create", service.CategoryCreate)
	// 分类修改
	group.PUT("/category-update", service.CategoryUpdate)
	// 分类删除
	group.DELETE("/category-delete", service.CategoryDelete)
	user := middleware.AuthUserCheck()
	// 用户私有方法
	users := engine.Group("", user)
	users.POST("/submit", service.Submit)

	// 下载速率选择器
	engine.LoadHTMLFiles("./rate.html")
	engine.GET("/rate", func(context *gin.Context) {
		context.HTML(http.StatusOK, "rate.html", nil)
	})
	engine.POST("/rates", service.Rate)

	// 自定义下载速度
	engine.GET("/speed", limiter.SpeedLim)

	engine.GET("/rates", service.Demo)

	// 初始化数据库
	engine.POST("/data", service.Databases)
	return engine

}
