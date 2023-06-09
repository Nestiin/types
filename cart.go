package types

import (
	"github.com/jinzhu/gorm"
)

// Cart req model
type AddToCart struct {
	gorm.Model
	OrderCode string `json:"orderCode" gorm:"type:varchar(50);column:order_code"`
	ProductId int    `json:"productId" gorm:"type:int;column:product_id"`
	UserID    int     `json:"userId" gorm:"type:int;column:user_id"`
	VariantId int     `json:"variantId" gorm:"type:int;column:variant_id"`
	Quantity  int     `json:"quantity" gorm:"type:int;column:quantity"`
	Price     int     `json:"price" gorm:"type:int;column:price"`
	ImageURLs []Image `json:"imageUrl" gorm:"-"`
}

func (AddToCart) TableName() string {
	return "cart_info"
}

type ProductDetails struct {
	Quantity int     `json:"quantity" validate:"required"`
	Price    int     `json:"price"`
	ImageURL []Image `json:"imageUrl" validate:"required"`
}
