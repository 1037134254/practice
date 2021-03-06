package models

import "gorm/gorm"

type ProblemCategory struct {
	ProblemId     uint           `gorm:"column:problem_id;type:int(11);" json:"problem_id"`           // 问题的ID
	CategoryId    uint           `gorm:"column:category_id;type:int(11);" json:"category_id"`         // 分类的ID
	CategoryBasic *CategoryBasic `gorm:"foreignKey:id;references:category_id;" json:"category_basic"` // 关联分类的基础信息表
	gorm.Model
}

func (table *ProblemCategory) TableName() string {
	return "problem_category"
}
