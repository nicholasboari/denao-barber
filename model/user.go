package model

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       string `gorm:"type:uuid;primary_key"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
