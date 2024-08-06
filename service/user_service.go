package service

import (
	"errors"
	"task-management/model"
	"task-management/repository"
	"task-management/utils"

	"golang.org/x/crypto/bcrypt"
)

type User = model.User

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{userRepo: repo}
}
func (us *UserService) Register(email, password, name string) error {
	existingUser, _ := us.userRepo.GetUserByEmail(email)

	if existingUser.ID != 0 {
		return errors.New("user already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("invalid password")
	}

	user := User{
		Email:    email,
		Password: string(hashedPassword),
		Name:     name,
	}

	us.userRepo.CreateUser(user)
	return nil
}

func (us *UserService) Login(email, password string) (User, error) {
	user, err := us.userRepo.GetUserByEmail(email)
	if err != nil {
		return User{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return User{}, errors.New("invalid credentials")
	}

	return user, nil
}

func (us *UserService) Profile(token string) (User, error) {

	checkToken, err := utils.VerifyToken(token)
	if err != nil {
		return User{}, errors.New("sorry, you don't have access for this feature")
	}
	email, ok := checkToken["email"].(string)
	if !ok {
		return User{}, errors.New("sorry, you don't have access for this feature")
	}
	user, err := us.userRepo.GetUserByEmail(email)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
