package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title    string    `json:"title" gorm:"type:varchar(255)"`
	Content  string    `json:"content" gorm:"type:text"`
	UserId   uint      `json:"user_id" gorm:"not null;index"`
	User     User      `json:"user" gorm:"references:ID"`
	Comments []Comment `json:"comments" gorm:"foreignKey:ID"`
}
