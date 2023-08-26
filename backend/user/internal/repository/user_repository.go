package repository

import (
	"errors"
	"myhealth/backend/user/internal/model"
	"sync"
)

var (
	users     map[string]model.User
	usersLock sync.RWMutex
)

func init() {
	users = make(map[string]model.User)
}

func SaveUser(user model.User) error {
	usersLock.Lock()
	defer usersLock.Unlock()

	if _, exists := users[user.ID]; exists {
		return errors.New("user already exists!")
	}

	users[user.ID] = user
	return nil
}

func GetAllUsers() []model.User {
	usersLock.RLock()
	defer usersLock.RUnlock()

	usersList := make([]model.User, 0, len(users))
	for _, user := range users {
		usersList = append(usersList, user)
	}
	return usersList
}
