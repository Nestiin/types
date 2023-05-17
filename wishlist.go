package types

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type WishListReq struct {
	gorm.Model
	UserID  int            `json:"user_id" validate:"required"`
	Product postgres.Jsonb `json:"product" validate:"required"`
}

func (WishListReq) TableName() string {
	return "wishlist_info"
}
