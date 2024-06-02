package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Haircut struct {
	Date       time.Time `json:"date"`
	NameClient string    `json:"client_name"`
	ID         uuid.UUID `json:"id" gorm:"type:uuid"`
	Price      float64   `json:"price"`
}

func (u *Haircut) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return nil
}
