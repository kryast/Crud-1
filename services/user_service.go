package services

import (
	"github.com/kryast/Crud-1.git/models"
	"github.com/kryast/Crud-1.git/repositories"
)

type UserService interface {
	CreateUser(user *models.User) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(r repositories.UserRepository) UserService {
	return &userService{r}
}

func (s *userService) CreateUser(user *models.User) error {
	return s.repo.Create(user)
}
