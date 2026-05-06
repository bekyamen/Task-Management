package repository

import (
	"user-service/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	UpdateUser(user *models.User) error
	FindByResetToken(tokenHash string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) FindByResetToken(tokenHash string) (*models.User, error) {
	var user models.User
	err := r.db.Where("reset_token_hash = ?", tokenHash).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
