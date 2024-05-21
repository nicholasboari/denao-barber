package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
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
func (u *UserServiceImpl) Create(user *model.User) error {
	err := u.Validate.Struct(user)
	if err != nil {
		return err
	}
	return u.UserRepository.Save(user)
}

// GetUserByID implements UserService
func (u *UserServiceImpl) GetUserByID(ID uuid.UUID) (*model.User, error) {
	return u.UserRepository.GetUserByID(ID)
}
