package repository

import "github.com/nicholasboari/denao-barber/model"

type HaircutRepository interface {
	Save(haircut *model.Haircut)
}
