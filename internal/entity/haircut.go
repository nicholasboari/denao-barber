package entity

import (
	"time"

	"github.com/nicholasboari/denao-barber/pkg/entity"
)

type Haircut struct {
	ID         entity.ID `json:"id"`
	NameClient string    `json:"nameClient"`
	Price      float64   `json:"price"`
	Date       time.Time `json:"date"`
}

func NewHaircut(nameClient string, price float64, dateStr string) (*Haircut, error) {
	date, err := time.Parse(time.DateTime, dateStr)
	if err != nil {
		return nil, err
	}
	return &Haircut{
		ID:         entity.NewID(),
		NameClient: nameClient,
		Price:      price,
		Date:       date,
	}, nil
}
