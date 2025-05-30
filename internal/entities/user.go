package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Username  string         `json:"username" gorm:"size:255;not null;unique"`
	Email     string         `json:"email" gorm:"size:255;not null;unique"`
	Password  string         `json:"-" gorm:"size:255;not null"`
	CreatedAt *time.Time     `json:"created_at" gorm:"default:null"`
	UpdatedAt *time.Time     `json:"updated_at" gorm:"default:null"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (User) TableName() string {
	return "users"
}
