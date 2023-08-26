package service

import (
	"myhealth/backend/user/internal/model"
	"myhealth/backend/user/internal/repository"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) AddUser(user model.User) error {
	return repository.SaveUser(user)
}

func (s *UserService) GetAllUsers() []model.User {
	return repository.GetAllUsers()
}
