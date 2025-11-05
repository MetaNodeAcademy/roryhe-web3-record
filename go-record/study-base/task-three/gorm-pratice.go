package main

import (
	"context"
	"fmt"
	"strconv"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	gorm.Model
	Name      string `gorm:"size:255"`
	Age       int    `gorm:"default:0"`
	PostCount int    `gorm:"default:0"`
	PostList  []Post
}

type Post struct {
	gorm.Model
	Title         string `gorm:"size:255"`
	Content       string `gorm:"type:text"`
	UserId        uint   `gorm:"not null"`
	CommentStatus int    `gorm:"default:0"` //0无评论，1有评论
	CommentList   []Comment
}

type Comment struct {
	gorm.Model
	PostId  uint   `gorm:"not null"`
	UserId  uint   `gorm:"not null"`
	Content string `gorm:"type:text"`
}

func create(db *gorm.DB) {
	errCreate := db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if errCreate != nil {
		fmt.Printf("create-error=%v\n", errCreate)
	}

	user := User{
		Name: "blog",
		Age:  20,
	}

	result := gorm.WithResult()
	ctx := context.Background()
	err := gorm.G[User](db, result).Create(ctx, &user)
	fmt.Printf("create-error=%v\n", err)

	post := Post{
		Title:   "文章",
		Content: "文章内容",
		UserId:  user.ID,
	}

	resultPost := gorm.WithResult()
	errPost := gorm.G[Post](db, resultPost).Create(ctx, &post)
	fmt.Printf("create-error=%v\n", errPost)

	for i := 0; i < 3; i++ {
		comment := Comment{
			PostId:  post.ID,
			UserId:  user.ID,
			Content: strconv.Itoa(i),
		}

		resultComment := gorm.WithResult()
		errComment := gorm.G[Comment](db, resultComment).Create(ctx, &comment)

		fmt.Printf("create-error=%v\n", errComment)
	}

	var comment Comment
	db.Model(&Comment{}).Order("id desc").First(&comment)
	db.Delete(&comment)
}

func queryPostAndCommentByUser(db *gorm.DB) *User {
	var user User
	err := db.Preload("PostList.CommentList").
		Where("name = ?", "blog").First(&user).Error
	if err != nil {
		fmt.Printf("failed to load user %+v\n", err)
	}

	fmt.Printf("load user %+v\n", user.PostList)
	return &user
}

func queryPostByMaxComments(db *gorm.DB) {
	type queryResult struct {
		PostId       uint
		CommentCount int
	}

	var result queryResult

	err := db.Model(&Comment{}).
		Select("post_id, count(*) as comment_count").
		Group("post_id").
		Order("comment_count").
		Limit(10).Scan(&result).Error
	if err != nil {
		fmt.Printf("failed to load result %+v\n", err)
	}

	fmt.Printf("load result %+v\n", result)

	var post Post
	errPost := db.Preload("CommentList").
		First(&post, result.PostId).Error

	if err != nil {
		fmt.Printf("failed to load post %+v\n", errPost)
	}

	fmt.Printf("load post %+v\n", post)
}

func (post *Post) AfterCreate(db *gorm.DB) (err error) {
	err = db.Model(&User{}).Where("id=?", post.UserId).
		UpdateColumn("post_count", gorm.Expr("post_count + ?", 1)).Error
	if err != nil {
		fmt.Printf("failed to update user %+v\n", err)
	}
	return
}

func (comment *Comment) AfterDelete(db *gorm.DB) (err error) {

	var count int64
	err = db.Model(&Comment{}).Where("post_id=?", comment.PostId).
		Count(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		err = db.Model(&Post{}).Where("id=?", comment.PostId).
			UpdateColumn("comment_status", 0).Error
	} else {
		err = db.Model(&Post{}).Where("id=?", comment.PostId).
			UpdateColumn("comment_status", 1).Error
	}

	if err != nil {
		return err
	}

	return
}

func main() {
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	fmt.Printf("Open-error=%v\n", err)
	create(db)
	//queryPostAndCommentByUser(db)
	//queryPostByMaxComments(db)
}
