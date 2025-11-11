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
	//TODO implement me
	panic("implement me")
}

func (c commentService) GetCommentByPostID(id uint) ([]models.Comment, error) {
	//TODO implement me
	panic("implement me")
}
