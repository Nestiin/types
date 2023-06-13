package types

import "github.com/jinzhu/gorm"

type Margin struct {
	gorm.Model
	Code       string `json:"code" gorm:"type:boolean;column:code"`
	Percentage int    `json:"percentage" gorm:"type:integer;column:percentage"`
	MerchantID int    `json:"merchantID" gorm:"type:integer;column:merchant_id"`
	Base       string `json:"base" gorm:"type:varchar;column:base"`
	Status     string `json:"status" gorm:"type:varchar;column:status"`
}

func (Margin) TableName() string {
	return "margin_info"
}