package service

import (
	"github.com/childelins/go-skeleton/app/model"
	repo "github.com/childelins/go-skeleton/app/repository"
)

type UserService struct {
	userRepo *repo.UserRepository
}

func NewUserService(u *repo.UserRepository) *UserService {
	return &UserService{
		userRepo: u,
	}
}

func (u *UserService) Create(user *model.User) bool {
	return u.userRepo.Create(user)
}

func (u *UserService) List(page, limit int) []*model.User {
	return u.userRepo.List(page, limit)
}
