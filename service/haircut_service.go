package service

import (
	"github.com/nicholasboari/denao-barber/model"
)

type HaircutService interface {
	Create(haircut *model.Haircut)
}
