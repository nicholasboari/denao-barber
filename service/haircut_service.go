package service

import (
	"github.com/google/uuid"
	"github.com/nicholasboari/denao-barber/model"
)

type HaircutService interface {
	Create(haircut *model.Haircut)
	GetHaircutByID(ID uuid.UUID) (*model.Haircut, error)
	GetAllHaircuts() ([]*model.Haircut, error)
}
