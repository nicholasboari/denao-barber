package repository

import "github.com/nicholasboari/denao-barber/model"

type UserRepository interface {
	Save(user *model.User)
}
