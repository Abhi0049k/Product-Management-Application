package models

import (
	"time"
)

// type User struct {
// 	gorm.Model
// 	ID        int
// 	Name      string
// 	Mobile    string
// 	Latitude  float64
// 	Longitude float64
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }

type User struct {
	ID        int       `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	Mobile    string    `json:"mobile"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
