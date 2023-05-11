package types

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

// Cart req model
type AddToCart struct {
	gorm.Model
	OrderCode string         `json:"orderCode" gorm:"type:varchar(50);column:order_code"`
	Product   postgres.Jsonb `json:"product" gorm:"type:json;column:product"`
	UserID    int            `json:"userId" gorm:"type:int;column:user_id"`
}

type ProductDetails struct {
	ProductID int    `json:"productId" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required"`
	Price     int    `json:"price"`
	ImageURL  string `json:"imageUrl" validate:"required"`
}
