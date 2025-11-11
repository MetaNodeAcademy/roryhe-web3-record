package service

import (
	"errors"
	"fmt"

	"github.com/rory7/task-four/database"
	"github.com/rory7/task-four/models"
	"gorm.io/gorm"
)

// UserService 用户服务接口
type UserService interface {
	// CreateUser 创建用户
	CreateUser(user *models.User) error
	// GetUserByID 根据ID获取用户
	GetUserByID(id uint) (*models.User, error)
	// GetUserByEmail 根据邮箱获取用户
	GetUserByEmail(email string) (*models.User, error)
	// GetAllUsers 获取所有用户
	GetAllUsers() ([]models.User, error)
	// UpdateUser 更新用户
	UpdateUser(id uint, user *models.User) error
	// DeleteUser 删除用户（软删除）
	DeleteUser(id uint) error
	// GetUserByName 通过用户名查找
	GetUserByName(name string) (*models.User, error)
}

type userService struct {
	db *gorm.DB
}

// NewUserService 创建用户服务实例
func NewUserService() UserService {
	return &userService{
		db: database.GetDB(),
	}
}

// CreateUser 创建用户
func (s *userService) CreateUser(user *models.User) error {
	// 检查邮箱是否已存在
	var existingUser models.User
	if err := s.db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return errors.New("邮箱已存在")
	}

	// 创建用户
	if err := s.db.Create(user).Error; err != nil {
		return fmt.Errorf("创建用户失败: %w", err)
	}

	return nil
}

func (s *userService) GetUserByName(username string) (*models.User, error) {
	var user models.User
	if err := s.db.Where("name = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUserByID 根据ID获取用户
func (s *userService) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := s.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, fmt.Errorf("查询用户失败: %w", err)
	}
	return &user, nil
}

// GetUserByEmail 根据邮箱获取用户
func (s *userService) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, fmt.Errorf("查询用户失败: %w", err)
	}
	return &user, nil
}

// GetAllUsers 获取所有用户
func (s *userService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := s.db.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("查询用户列表失败: %w", err)
	}
	return users, nil
}

// UpdateUser 更新用户
func (s *userService) UpdateUser(id uint, user *models.User) error {
	// 检查用户是否存在
	var existingUser models.User
	if err := s.db.First(&existingUser, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("用户不存在")
		}
		return fmt.Errorf("查询用户失败: %w", err)
	}

	// 如果更新邮箱，检查新邮箱是否已被其他用户使用
	if user.Email != "" && user.Email != existingUser.Email {
		var emailUser models.User
		if err := s.db.Where("email = ? AND id != ?", user.Email, id).First(&emailUser).Error; err == nil {
			return errors.New("邮箱已被其他用户使用")
		}
	}

	// 更新用户
	if err := s.db.Model(&existingUser).Updates(user).Error; err != nil {
		return fmt.Errorf("更新用户失败: %w", err)
	}

	return nil
}

// DeleteUser 删除用户（软删除）
func (s *userService) DeleteUser(id uint) error {
	// 检查用户是否存在
	var user models.User
	if err := s.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("用户不存在")
		}
		return fmt.Errorf("查询用户失败: %w", err)
	}

	// 软删除
	if err := s.db.Delete(&user).Error; err != nil {
		return fmt.Errorf("删除用户失败: %w", err)
	}

	return nil
}
