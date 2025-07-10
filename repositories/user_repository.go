package repositories

import (
	"github.com/kryast/Crud-1.git/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}
