package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/nicholasboari/denao-barber/helper"
	"github.com/nicholasboari/denao-barber/model"
	"github.com/nicholasboari/denao-barber/repository"
)

type HaircutServiceImpl struct {
	HaircutRepository repository.HaircutRepository
	Validate          *validator.Validate
}

func NewHaircutServiceImpl(haircutRepository repository.HaircutRepository, validate *validator.Validate) HaircutService {
	return &HaircutServiceImpl{HaircutRepository: haircutRepository,
		Validate: validate}
}

// Create implements HaircutService
func (h *HaircutServiceImpl) Create(haircut *model.Haircut) {
	err := h.Validate.Struct(haircut)
	helper.ErrorPanic(err)
	h.HaircutRepository.Save(haircut)
}