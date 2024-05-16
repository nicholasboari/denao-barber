package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/nicholasboari/denao-barber/helper"
	"github.com/nicholasboari/denao-barber/model"
	"github.com/nicholasboari/denao-barber/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewUserServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{UserRepository: userRepository, Validate: validate}
}

// Create implements UserService
func (u *UserServiceImpl) Create(user *model.User) {
	err := u.Validate.Struct(user)
	helper.ErrorPanic(err)
	u.UserRepository.Save(user)
}
