package repository

import (
	"github.com/nicholasboari/denao-barber/helper"
	"github.com/nicholasboari/denao-barber/model"
	"gorm.io/gorm"
)

type HaircutRepositoryImpl struct {
	Db *gorm.DB
}

func NewHaircutRepositoryImpl(Db *gorm.DB) HaircutRepository {
	return &HaircutRepositoryImpl{Db: Db}
}

// Save implements HaircutRepository
func (h *HaircutRepositoryImpl) Save(haircut *model.Haircut) {
	result := h.Db.Create(&haircut)
	helper.ErrorPanic(result.Error)
}
