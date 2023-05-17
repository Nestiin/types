package types

import "github.com/jinzhu/gorm/dialects/postgres"

type WishListReq struct {
	UserID  int            `json:"user_id" validate:"required"`
	Product postgres.Jsonb `json:"product" validate:"required"`
}
