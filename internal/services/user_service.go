package services

import (
	"errors"
	"github.com/antibomberman/junior_test/internal/models"
	"github.com/antibomberman/junior_test/internal/repository"
)

type UserService interface {
	Create(models.UserCreate) (models.User, error)
	GetById(string) (models.User, error)
	All() ([]models.User, error)
	Update(string, models.UserUpdate) (models.User, error)
	Delete(string) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}
func (u *userService) Create(data models.UserCreate) (models.User, error) {
	id, err := u.userRepo.CreateUser(data)
	if err != nil {
		return models.User{}, err
	}
	user, err := u.userRepo.GetUserByID(string(id))

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
func (u *userService) GetById(userId string) (models.User, error) {
	user, err := u.userRepo.GetUserByID(userId)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
func (u *userService) All() ([]models.User, error) {
	users, err := u.userRepo.GetAllUsers(0, 0, "")
	if err != nil {
		return []models.User{}, err
	}
	return users, nil
}
func (u *userService) Update(userId string, data models.UserUpdate) (models.User, error) {
	has, err := u.userRepo.HasUser(userId)
	if err != nil {
		return models.User{}, err
	}
	if !has {
		return models.User{}, errors.New("not found")
	}
	_, err = u.userRepo.UpdateUser(userId, data)
	if err != nil {
		return models.User{}, err
	}
	user, err := u.userRepo.GetUserByID(userId)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
func (u *userService) Delete(userId string) error {
	err := u.userRepo.DeleteUser(userId)
	if err != nil {
		return err
	}
	return nil
}
