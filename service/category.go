package service

import (
	"example.com/m/v2/defaults"
	"example.com/m/v2/helper"
	"example.com/m/v2/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetCategoryList
// @Tags 管理员私有方法
// @Summary 分类列表
// @Param authorization header string true "authorization"
// @Param page query int false "page"
// @Param size query int false "size"
// @Param keyword query string false "keyword"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /category-list [get]
func GetCategoryList(c *gin.Context) {
	size, _ := strconv.Atoi(c.DefaultQuery("size", defaults.DefaultSize))
	page, err := strconv.Atoi(c.DefaultQuery("size", defaults.DefaultPage))
	if err != nil {
		log.Printf("Get GetCategoryList Error:" + err.Error())
	}
	page = (page - 1) * size
	var count int64
	keyword := c.Query("keyword")
	categoryList := make([]*models.CategoryBasic, 0)
	err = models.DB.Debug().Model(new(models.CategoryBasic)).Where("name like ?", "%"+keyword+"%").
		Count(&count).Offset(page).Limit(size).Find(&categoryList).Error
	if err != nil {
		log.Printf("Get CList Error:" + err.Error())
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "获取分类列表失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"date": map[string]interface{}{
			"list":  categoryList,
			"count": count,
		},
	})

}

// CategoryCreate
// @Tags 管理员私有方法
// @Summary 创建分类
// @Param authorization header string true "authorization"
// @Param name formData string true "name"
// @Param parentId formData int false "parentId"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /category-create [post]
func CategoryCreate(c *gin.Context) {
	name := c.PostForm("name")
	parentId, _ := strconv.Atoi(c.PostForm("parentId"))

	category := &models.CategoryBasic{
		Identity: helper.UUid(),
		Name:     name,
		ParentId: parentId,
	}

	err := models.DB.Create(category).Debug().Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "创建分类失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "创建成功",
	})
	return
}

// CategoryUpdate
// @Tags 管理员私有方法
// @Summary 修改分类
// @Param authorization header string true "authorization"
// @Param identity formData string true "identity"
// @Param name formData string true "name"
// @Param parentId formData int false "parentId"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /category-update [put]
func CategoryUpdate(c *gin.Context) {
	identity := c.PostForm("identity")
	name := c.PostForm("name")
	parentId, _ := strconv.Atoi(c.PostForm("parentId"))
	category := &models.CategoryBasic{
		Identity: identity,
		Name:     name,
		ParentId: parentId,
	}
	if identity == "" || name == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不正确",
		})
		return
	}
	err := models.DB.Model(new(models.CategoryBasic)).Where("identity = ?", identity).Updates(category).Error
	if err != nil {
		log.Println("更新失败：" + err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "CategoryUpdate Error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "更新成功",
	})
	return
}

// CategoryDelete
// @Tags 管理员私有方法
// @Summary 删除分类
// @Param authorization header string true "authorization"
// @Param identity query string false "identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /category-delete [delete]
func CategoryDelete(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不正确",
		})
		return
	}
	var cnt int64
	err := models.DB.Model(new(models.ProblemCategory)).
		Where("category_id = (select id from category_basic where identity = ? limit 1)", identity).
		Count(&cnt).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "获取分类关联的问题失败",
		})
		return
	}
	if cnt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "该分类下面已存在问题，不可删除",
		})
		return
	}
	var i int64
	models.DB.Model(new(models.CategoryBasic)).Where("identity = ? ", identity).Count(&i)
	if i <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "问题不存在",
		})
		return
	}
	err = models.DB.Model(new(models.CategoryBasic)).Where("identity = ? ", identity).Delete(new(models.CategoryBasic)).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "删除失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "删除成功",
	})
	return
}
