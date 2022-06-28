package models

type UserBasic struct {
	// 用户唯一标识
	Identity string `gorm:"colum:identity;type:varchar(36);" json:"identity"`
	Name     string `gorm:"cloum:name;type:varchar(100);" json:"name"`
	Password string `gorm:"cloum:password;type:varchar(100);" json:"password"`
	Phone    string `gorm:"cloum:phone;type:varchar(100);" json:"phone"`
	// 完成问题的个数
	FinisProblemNum int64  `gorm:"cloum:finisProblemNum;type:int(11);" json:"finisProblemNum"`
	Mail            string `gorm:"cloum:mail;type:varchar(100);" json:"mail"`
	// 提交次数
	SubmitNum int64 `gorm:"cloum:submitNum;type:int(11);" json:"submitNum"`
	// 0不是管理员 1 是管理员
	IsAdmin int `gorm:"cloum:isAdmin;type:tinyint(1);" json:"isAdmin"`
	// 通过的次数
	PassNum int64 `gorm:"column:pass_num;type:int(11);" json:"passNum""`
}

func (table UserBasic) TableName() string {
	return "user_basic"
}
