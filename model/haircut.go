package model

import "time"

type Haircut struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	NameClient string    `json:"client_name"`
	Price      float64   `json:"price"`
	Date       time.Time `json:"date"`
}
