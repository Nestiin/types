package types

import "github.com/jinzhu/gorm"

type ProductVariantInfo struct {
	gorm.Model
	ProductId uint    `json:"productId,omitempty" gorm:"type:bigint;column:product_id"`
	Size      string  `json:"size,omitempty" gorm:"type:varchar(250);column:size"`
	Color     string  `json:"color,omitempty" gorm:"type:varchar(250);column:color"`
	Status    string  `json:"status,omitempty" gorm:"type:varchar(20);column:status;default:PENDING_APPROVAL"`
	Quantity  int     `json:"quantity,omitempty" gorm:"type:int4;column:quantity"`
	UnitPrice float64 `json:"price,omitempty" gorm:"type:float4;column:unit_price"`
	Discount  int     `json:"discount,omitempty" gorm:"type:int;column:discount"`
	Height    string  `json:"height" gorm:"type:varchar(50);column:height"`
	Length    string  `json:"length" gorm:"type:varchar(50);column:length"`
	Width     string  `json:"width" gorm:"type:varchar(50);column:width"`
	Weight    string  `json:"weight" gorm:"type:varchar(50);column:weight"`
	Comment   string  `json:"comment" gorm:"type:varchar(250);column:review_comment"`
	IsDefault bool    `json:"isDefault,omitempty" gorm:"type:boolean;column:is_default"`
}

func (ProductVariantInfo) TableName() string {
	return "product_variant_info"
}
