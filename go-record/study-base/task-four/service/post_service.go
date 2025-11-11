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
	if err := p.db.Create(post).Error; err != nil {
		return 0, err
	}

	return post.ID, nil
}

func (p *postService) GetALLPosts() ([]models.Post, error) {
	var posts []models.Post
	if err := p.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *postService) GetPostById(postId uint) (*models.Post, error) {
	var post models.Post
	if err := p.db.Where("id = ?", postId).First(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func (p *postService) DeletePostById(postId uint) error {
	if err := p.db.Delete(&models.Post{}, postId).Error; err != nil {
		return err
	}
	return nil
}

func (p *postService) UpdatePost(post *models.Post) error {
	if err := p.db.Save(post).Error; err != nil {
		return err
	}

	return nil
}
