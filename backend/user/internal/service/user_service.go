package service

import (
	"myhealth/backend/user/internal/model"
	"myhealth/backend/user/internal/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) CreateUser(user model.User) error {
	return s.userRepository.SaveUser(user)
}

func (s *UserService) GetUsers() ([]model.User, error) {
	return s.userRepository.GetAllUsers()
}
