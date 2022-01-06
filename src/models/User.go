package models

import "time"

type User struct {
	ID        uint64    `gorm:"autoIncrement;primaryKey" json:"id,omitempty"`
	Name      string    `gorm:"type:varchar(50);not null" json:"name,omitempty"`
	Nick      string    `gorm:"type:varchar(50);not null;uniqueIndex" json:"nick,omitempty"`
	Email     string    `gorm:"type:varchar(50);not null;uniqueIndex" json:"email,omitempty"`
	Password  string    `gorm:"type:varchar(20);not null" json:"password,omitempty"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt,omitempty"`
}
