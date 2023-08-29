package repository

import (
	"database/sql"
	"myhealth/backend/user/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) SaveUser(user model.User) error {
	_, err := r.db.Exec("INSERT INTO users (id, username, email) VALUES (?, ?, ?)", user.ID, user.Username, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetAllUsers() ([]model.User, error) {
	rows, err := r.db.Query("SELECT id, username, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	usersList := make([]model.User, 0)
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			return nil, err
		}
		usersList = append(usersList, user)
	}
	return usersList, nil
}
