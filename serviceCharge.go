package types

import "github.com/jinzhu/gorm"

type ServiceCharge struct {
	gorm.Model
	Code              string `json:"code" gorm:"type:boolean;column:code"`
	Silver            bool   `json:"silver" gorm:"type:boolean;column:silver"`
	Gold              bool   `json:"gold" gorm:"type:boolean;column:gold"`
	Platinum          bool   `json:"platinum" gorm:"type:boolean;column:platinum"`
	ServiceCharge     int    `json:"serviceCharge" gorm:"type:integer;column:service_charge"`
	MerchantID        int    `json:"merchantID" gorm:"type:integer;column:merchant_id"`
	ProductTypeID     int    `json:"productTypeID" gorm:"type:integer;column:product_type_id"`
	ProductCategoryID int    `json:"productCategoryID" gorm:"type:integer;column:product_category_id"`
	ProductID         int    `json:"productID" gorm:"type:integer;column:product_id"`
	Status            bool   `json:"status" gorm:"type:varchar(50);column:status"`
}

func (ServiceCharge) TableName() string {
	return "service_charge_info"
}
