package service

import (
	"example.com/m/v2/defaults"
	"example.com/m/v2/models"
	"github.com/gin-gonic/gin"
	"gorm/gorm"
	"log"
	"net/http"
	"strconv"
)

// GetProblemList
// @Tags 问题
// @Summary 问题列表
// @Param page query int false "page"
// @Param size query int false "size"
// @Param keyword query string false "keyword"
// @Param category_identity query string false "category_identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /problem-list [get]
func GetProblemList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", defaults.DefaultPage))
	size, err := strconv.Atoi(c.DefaultQuery("page", defaults.DefaultSize))
	if err != nil {
		log.Printf("error", err)
		return
	}
	page = (page - 1) * size
	var count int64
	keyword := c.Query("keyword")
	categoryIdentity := c.Query("category_identity")
	list := make([]*models.ProblemBasic, 0)
	tx := models.GetProblemList(keyword, categoryIdentity)
	err = tx.Count(&count).Omit("content").Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"list":  list,
			"count": count,
		},
	})
}

// GetProblemDetail
// @Tags 问题
// @Summary 问题详情
// @Param identity query string false "problem_identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /problem-detail [get]
func GetProblemDetail(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "唯一标识符不能为空",
		})
		return
	}
	date := new(models.ProblemBasic)
	err := models.DB.Where("?", identity).Preload("ProblemCategory.CategoryBasic").First(&date).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "问题不存在",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get Problem Error " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  date,
	})
	return
}

// ProblemCreate
// @Tags 管理员私有方法
// @Summary 问题创建
// @Param identity query string false "problem_identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /problem-detail [get]
//func ProblemCreate(c *gin.Context) {
//	Title := c.PostForm("title")
//	Content := c.PostForm("content")
//	MaxRuntime, _ := strconv.Atoi(c.PostForm("max_runtime"))
//	MaxMen, _ := strconv.Atoi(c.PostForm("max_men"))
//	categoryIds := c.PostForm("category_ids")
//	TestCases := c.PostForm("test_case")
//
//	if Title == "" || Content == "" || len(categoryIds) == 0 || len(TestCase) == 0 {
//		c.JSON(http.StatusOK, gin.H{
//			"code": -1,
//			"msg":  "参数不能为空",
//		})
//		return
//	}
//	identity := helper.UUid()
//	date := &models.ProblemBasic{
//		Identity:   identity,
//		Title:      Title,
//		Content:    Content,
//		MaxRuntime: MaxRuntime,
//		MaxMen:     MaxMen,
//	}
//	// 处理分类
//	problemCategory := make([]*models.ProblemCategory, 0)
//	for _, id := range problemCategory {
//		problemCategory = append(problemCategory, &models.ProblemCategory{
//			ProblemId:  date.ID,
//			CategoryId: id.CategoryId,
//		})
//	}
//	date.ProblemCategory = problemCategory
//	// 处理测试用例
//	testCaseBasic := make([]*models.TestCase, 0)
//	for _, testCase := range TestCases {
//		caseMap := make(map[string]string)
//		json.Unmarshal([]byte(testCase), &caseMap)
//	}
//	err := models.DB.Create(&date).Error
//	if err != nil {
//		c.JSON(http.StatusOK, gin.H{
//			"code": -1,
//			"msg":  "创建失败" + err.Error(),
//		})
//		return
//	}
//	c.JSON(http.StatusOK, gin.H{
//		"code": http.StatusOK,
//		"date": map[string]interface{}{
//			"identity": date.Identity,
//		},
//	})
//}
