package entities

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Title    string `json:"title" gorm:"size:255;not null"`
	Director string `json:"director" gorm:"size:255;not null"`
	Year     int    `json:"year" gorm:"not null"`
	Plot     string `json:"plot" gorm:"type:text"`
	// ReleaseDate *time.Time     `json:"release_date" gorm:"default:null"`
	CreatedAt *time.Time     `json:"created_at" gorm:"default:null"`
	UpdatedAt *time.Time     `json:"updated_at" gorm:"default:null"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (Movie) TableName() string {
	return "movies"
}
