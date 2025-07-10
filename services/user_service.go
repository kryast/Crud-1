package services

import (
	"github.com/kryast/Crud-1.git/models"
	"github.com/kryast/Crud-1.git/repositories"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetUsers() ([]models.User, error)
	GetUserByID(id uint) (models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
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

func (s *userService) GetUsers() ([]models.User, error) {
	return s.repo.FindAll()
}

func (s *userService) GetUserByID(id uint) (models.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) UpdateUser(user *models.User) error {
	return s.repo.Update(user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}
