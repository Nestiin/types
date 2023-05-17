package types

type WishListReq struct {
	UserID  int       `json:"user_id" validate:"required"`
	Product []Product `json:"product" validate:"required"`
}
