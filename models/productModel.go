package models

import (
	"time"

	"github.com/lib/pq"
)

type Product struct {
	ProductID               int            `gorm:"primary_key" json:"product_id"`
	ProductName             string         `json:"product_name"`
	ProductDescription      string         `json:"product_description"`
	ProductImages           pq.StringArray `json:"product_images" gorm:"type:text[]"`
	ProductPrice            uint32         `json:"product_price"`
	CompressedProductImages pq.StringArray `json:"compressed_product_images" gorm:"type:text[]"`
	CreatedAt               time.Time      `json:"created_at"`
	UpdatedAt               time.Time      `json:"updated_at"`
}
