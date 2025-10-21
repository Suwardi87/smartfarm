package user

import (
	"log"
	"fmt"
    "smartfarm-api/internal/middleware"
	"smartfarm-api/pkg/utils"
)

type Service interface {
	Register(input User) (*User, error)
	Login(username, password string) (string, *User, error)
}

type userService struct {
	repo *userRepository
}


func NewService(repo *userRepository) userService {
	return userService{repo: repo}
}

func (s *userService) Register(input User) (*User, error) {
	hashed, err := utils.HashPassword(input.Password)
	if err != nil {
		log.Println("hash error:", err)
		return nil, err
	}
	input.Password = hashed

	if err := s.repo.Create(&input); err != nil {
		return nil, err
	}

	return &input, nil
}

func (s userService) Login(username, password string) (string, *User, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return "", nil, fmt.Errorf("user not found")
	}

	if !utils.CheckPassword(user.Password, password) {
		return "", nil, fmt.Errorf("wrong password")
	}

	token, err := middleware.GenerateJWT(user.ID, user.Username, user.Level)
	if err != nil {
		return "", nil, fmt.Errorf("failed to generate token")
	}

	return token, user, nil
}

func (u *userService) GetAllUser() ([]User, error) {
	return u.repo.GetAllUser()
}