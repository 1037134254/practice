package models

import (
	"fmt"
	gorm2 "gorm.io/gorm"
	"gorm/gorm"
)

type ProblemBasic struct {
	// 关联问题分类表
	ProblemCategory []*ProblemCategory `gorm:"foreignKey:problem_id;references:id" json:"problem_categories"`
	// 问题的唯一标识
	Identity   string `gorm:"column:identity;type:varchar(36);" json:"identity"`
	CategoryId string `gorm:"column:title;type:varchar(255);" json:"categoryId"`
	Title      string `gorm:"column:title;type:varchar(255);" json:"title"`
	Content    string `gorm:"column:content;type:varchar(255);" json:"content"`
	// 最大运行时间
	MaxRuntime int `gorm:"maxRuntime:mail;type:int(11);" json:"maxRuntime"`
	// 最大的运行内存
	MaxMen int `gorm:"cloum:maxMen;type:int(11);" json:"maxMen"`
	// 关联测试用例表
	TestCases []*TestCase `gorm:"foreignKey:problem_identity;references:identity;" json:"testCase"`
	// 通过次数
	PassNum int64 `gorm:"column:pass_num;type:int(11);" json:"passNum"`
	// 提交次数
	SubmitNum int64 `gorm:"column:submit_num;type:int(11);" json:"submitNum"`
	gorm.Model
}

func (table *ProblemBasic) TableName() string {
	return "problem_basic"
}
func GetProblem() {
	date := make([]*ProblemBasic, 0)
	DB.Find(&date)
	for _, problem := range date {
		fmt.Printf("problem:%v\n", problem)
	}
}

func GetProblemList(keyword, categoryIdentity string) *gorm2.DB {
	tx := DB.Model(new(ProblemBasic)).Preload("ProblemCategory").Preload("ProblemCategory.CategoryBasic").
		Where("title like ? OR content like ?", "%"+keyword+"%", "%"+keyword+"%")
	if categoryIdentity != "" {
		tx.Joins("RIGHT JION problem_category pc on pc.problem_id = problem_basid.id").
			Where("pc.category_id=(SELECT cb.id FROM category_basic cb WHERE cb.identity =?)", categoryIdentity)
	}
	return tx
}
