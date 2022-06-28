package test

import (
	"example.com/m/v2/models"
	"fmt"
	"gorm.io/gorm"
	"gorm/mysql"
	"strconv"
	"testing"
)

func TestDB_test(t *testing.T) {
	open, err := gorm.Open(mysql.Open("root:root1234@(127.0.0.1:13306)/OnlinePractice?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		return
	}
	open.AutoMigrate(models.CategoryBasic{})
	open.AutoMigrate(models.ProblemBasic{})
	open.AutoMigrate(models.ProblemCategory{})
	open.AutoMigrate(models.SubmitBasic{})
	open.AutoMigrate(models.UserBasic{})
	open.AutoMigrate(models.TestCase{})
}
func TestDemo(t *testing.T) {
	var i int64
	models.DB.Model(new(models.CategoryBasic)).
		Where("identity = ? ", "ccb35773-9c8f-482e-9363-03a77c81eb1c3").Count(&i)
	fmt.Printf(strconv.FormatInt(i, 10))
}
