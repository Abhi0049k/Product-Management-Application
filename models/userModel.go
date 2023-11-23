package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int
	Name      string
	Mobile    string
	Latitude  float64
	Longitude float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
