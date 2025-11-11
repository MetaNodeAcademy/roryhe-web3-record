package service

import (
	"github.com/rory7/task-four/database"
	"github.com/rory7/task-four/models"
	"gorm.io/gorm"
)

type CommentService interface {
	CreateComment(comment *models.Comment) (uint, error)
	GetCommentByPostID(id uint) ([]models.Comment, error)
}

type commentService struct {
	db *gorm.DB
}

func NewCommentService() CommentService {
	return &commentService{
		db: database.GetDB(),
	}
}

func (c commentService) CreateComment(comment *models.Comment) (uint, error) {
	if err := c.db.Create(&comment).Error; err != nil {
		return 0, err
	}

	return comment.ID, nil
}

func (c commentService) GetCommentByPostID(id uint) ([]models.Comment, error) {
	var comments []models.Comment
	if err := c.db.Where("post_id = ?", id).Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}
