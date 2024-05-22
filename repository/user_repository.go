package repository

import (
	"github.com/google/uuid"
	"github.com/nicholasboari/denao-barber/model"
)

type UserRepository interface {
	Save(user *model.User) error
	GetUserByID(ID uuid.UUID) (*model.User, error)
	Delete(ID uuid.UUID) error
	Update(user *model.User) (*model.User, error)
}
