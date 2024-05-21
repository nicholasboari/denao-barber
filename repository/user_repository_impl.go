package repository

import (
	"github.com/google/uuid"
	"github.com/nicholasboari/denao-barber/model"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}

// Save implements UserRepository
func (u *UserRepositoryImpl) Save(user *model.User) error {
	return u.Db.Create(&user).Error
}

// GetUserByID implements UserRepository
func (u *UserRepositoryImpl) GetUserByID(ID uuid.UUID) (*model.User, error) {
	var result model.User
	err := u.Db.Model(model.User{ID: ID}).First(&result).Error
	return &result, err
}

// Delete implements UserRepository
func (u *UserRepositoryImpl) Delete(ID uuid.UUID) error {
	user, err := u.GetUserByID(ID)
	if err != nil {
		return err
	}
	if err := u.Db.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
