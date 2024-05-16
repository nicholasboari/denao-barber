package repository

import (
	"github.com/nicholasboari/denao-barber/helper"
	"github.com/nicholasboari/denao-barber/model"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}

// Save implements UserRepository
func (u *UserRepositoryImpl) Save(user *model.User) {
	result := u.Db.Create(&user)
	helper.ErrorPanic(result.Error)
}
