package repository

import (
	"github.com/google/uuid"
	"github.com/nicholasboari/denao-barber/model"
)

type HaircutRepository interface {
	Save(haircut *model.Haircut)
	Delete(ID uuid.UUID) error
	GetHaircutByID(ID uuid.UUID) (*model.Haircut, error)
	GetAllHaircuts() ([]*model.Haircut, error)
}
