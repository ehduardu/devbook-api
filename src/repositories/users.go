package repositories

import (
	"api/src/models"
	"fmt"

	"gorm.io/gorm"
)

type Users struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *Users {
	return &Users{db}
}

func (repository Users) Create(user models.User) (models.User, error) {
	err := repository.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repository Users) List(nameOrNick string) ([]models.User, error) {
	var users []models.User

	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) //%nameOrNick%
	err := repository.db.Select("ID", "name", "nick", "email").Where("name ILIKE ?", nameOrNick).Or("nick LIKE ?", nameOrNick).Find(&users).Error

	return users, err
}
