package repository

import (
	"github.com/google/uuid"
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

// GetHaircutByID implements HaircutRepository
func (h *HaircutRepositoryImpl) GetHaircutByID(ID uuid.UUID) (*model.Haircut, error) {
	var result model.Haircut
	err := h.Db.Model(model.Haircut{ID: ID}).First(&result).Error
	return &result, err
}

// GetAllHaircuts implements HaircutRepository
func (h *HaircutRepositoryImpl) GetAllHaircuts() ([]*model.Haircut, error) {
	var result []*model.Haircut
	err := h.Db.Find(&result).Error
	return result, err
}
