package service

import "github.com/nicholasboari/denao-barber/model"

type UserService interface {
	Create(user *model.User)
}
