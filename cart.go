package types

import "github.com/jinzhu/gorm"

// Cart req model
type AddToCart struct {
	gorm.Model
	OrderCode string `json:"order_code" gorm:"type:varchar(50);column:order_code"`
	Product   string `json:"product" gorm:"type:json;column:product"`
	UserID    int    `json:"user_id" gorm:"type:int;column:user_id"`
}
