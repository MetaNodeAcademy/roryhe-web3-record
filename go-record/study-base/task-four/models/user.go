package models

import (
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	// Name 用户名
	Name string `json:"name" gorm:"size:255;not null"`
	// Email 邮箱
	Email string `json:"email" gorm:"size:255;uniqueIndex;not null"`
	//Password 密码
	Password string `json:"password" gorm:"size:255;not null"`
	Posts    []Post `json:"posts" gorm:"foreignKey:ID"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}
