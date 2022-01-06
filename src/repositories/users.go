package repositories

import (
	"api/src/models"

	"gorm.io/gorm"
)

type Users struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *Users {
	return &Users{db}
}

func (repository Users) Create(user models.User) error {

	err := repository.db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}
