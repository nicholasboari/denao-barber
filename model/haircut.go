package model

import "time"

type Haircut struct {
	Date       time.Time `json:"date"`
	NameClient string    `json:"client_name"`
	ID         uint      `json:"id" gorm:"primary_key"`
	Price      float64   `json:"price"`
}
