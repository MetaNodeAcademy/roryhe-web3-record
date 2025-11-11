package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `json:"content" gorm:"type:text;not null"`
	UserID  uint   `json:"user_id" gorm:"not null;index"`
	User    User   `json:"user" gorm:"references:ID"`
	PostID  uint   `json:"post_id" gorm:"not null;index"`
	Post    Post   `json:"post" gorm:"references:ID"`
}
