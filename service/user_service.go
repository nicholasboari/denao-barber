package service

import (
	"github.com/google/uuid"
	"github.com/nicholasboari/denao-barber/model"
)

type UserService interface {
	Create(user *model.User) error
	GetUserByID(ID uuid.UUID) (*model.User, error)
}
