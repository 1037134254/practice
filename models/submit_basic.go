package models

import (
	gorm2 "gorm.io/gorm"
	"gorm/gorm"
)

type SubmitBasic struct {
	// 唯一标识
	Identity string `gorm:"colum:identity;type:varchar(36);" json:"identity"`
	// 关联问题基础表
	ProblemBasic *ProblemBasic `gorm:"foreignKey:identity;references:problem_identity;" json:"problemBasic"`
	// 问题的唯一标识
	ProblemIdentity string `gorm:"colum:problemIdentity;type:varchar(255);" json:"problemIdentity"`
	// 关联用户的基础表
	UserBasic *UserBasic `gorm:"foreignKey:identity;references:user_identity;" json:"userBasic"`
	// 用户唯一标识
	UserIdentity string `gorm:"colum:title;type:varchar(255);" json:"userIdentity"`
	// 代码存放路径
	Path string `gorm:"colum:content;type:varchar(255);" json:"path"`
	// -1 判断，1 正确，2 错误，3 运行超时，4运行超内存,5编译错误
	Status int `gorm:"colum:status;type:tinyint(1);" json:"status"`
	gorm.Model
}

func (table *SubmitBasic) TableName() string {
	return "submit_basic"
}

func GetSubmitList(problemIdentity, userIdentity string, status int) *gorm2.DB {
	tx := DB.Model(new(SubmitBasic)).Preload("ProblemBasic", func(db *gorm2.DB) *gorm2.DB {
		return db.Omit("content")
	})
	//.Preload("UserBasic")
	if problemIdentity != "" || userIdentity != "" || status != 0 {
		tx.Where("problem_identity = ? ", problemIdentity)
	}
	if userIdentity != "" {
		tx.Where(" user_identity = ?", userIdentity)
	}
	if status != 0 {
		tx.Where(" status = ?", status)
	}
	return tx
}
