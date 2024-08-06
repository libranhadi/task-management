package repository

import (
	"errors"
	"task-management/model"
)

type User = model.User

type UserRepository struct {
	users  map[string]User
	nextID int
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users:  make(map[string]User),
		nextID: 1,
	}
}

func (repo *UserRepository) CreateUser(user User) {
	user.ID = repo.nextID
	repo.users[user.Email] = user
	repo.nextID++
}

func (repo *UserRepository) GetUserByEmail(email string) (User, error) {
	user, exists := repo.users[email]
	if !exists {
		return User{}, errors.New("account not found")
	}
	return user, nil
}
