package service

import (
	"myproject/internal/repo"
)

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

func (s UserService) GetInfoUserService() string {
	return "User Info Service"
}
