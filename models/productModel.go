package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductID               int
	ProductName             string
	ProductDescription      string
	ProductImages           []string
	ProductPrice            float64
	CompressedProductImages []string
	CreatedAt               time.Time
	UpdatedAt               time.Time
}
