package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID        uint64    `gorm:"autoIncrement;primaryKey" json:"id,omitempty"`
	Name      string    `gorm:"type:varchar(50);not null" json:"name,omitempty" validate:"required"`
	Nick      string    `gorm:"type:varchar(50);not null;uniqueIndex" json:"nick,omitempty" validate:"required"`
	Email     string    `gorm:"type:varchar(50);not null;uniqueIndex" json:"email,omitempty" validate:"required"`
	Password  string    `gorm:"type:varchar(20);not null" json:"password,omitempty" validate:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt,omitempty"`
}

func (user User) Check(stage string) error {
	var validate = validator.New()

	if stage == "register" {
		if err := validate.Struct(user); err != nil {
			return err
		}
	} else {
		if err := validate.StructExcept(user, "Password"); err != nil {
			return err
		}
	}
	return nil
}
