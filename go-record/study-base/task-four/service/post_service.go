package service

import (
	"github.com/rory7/task-four/database"
	"github.com/rory7/task-four/models"
	"gorm.io/gorm"
)

type PostService interface {
	CreatePost(post *models.Post) (uint, error)
	GetALLPosts() ([]models.Post, error)
	GetPostById(postId uint) (*models.Post, error)
	DeletePostById(postId uint) error
	UpdatePost(post *models.Post) error
}

type postService struct {
	db *gorm.DB
}

func NewPostService() PostService {
	return &postService{
		db: database.GetDB(),
	}
}

func (p *postService) CreatePost(post *models.Post) (uint, error) {
	//TODO implement me
	panic("implement me")
}

func (p *postService) GetALLPosts() ([]models.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p *postService) GetPostById(postId uint) (*models.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p *postService) DeletePostById(postId uint) error {
	//TODO implement me
	panic("implement me")
}

func (p *postService) UpdatePost(post *models.Post) error {
	//TODO implement me
	panic("implement me")
}
