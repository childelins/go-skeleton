package repository

import (
	"github.com/childelins/go-skeleton/app/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) Create(user *model.User) bool {
	u.db.Create(user)
	return user.ID > 0
}

func (u *UserRepository) List(page, limit int) []*model.User {
	var users []*model.User

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	offset := page - 1
	u.db.Offset(offset).Limit(limit).Find(&users)
	return users
}
