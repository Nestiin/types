package types

import "github.com/jinzhu/gorm"

type ProductType struct {
	gorm.Model
	Code        string `json:"code,omitempty" gorm:"type:varchar(50);column:code"`
	Name        string `json:"name,omitempty" gorm:"type:varchar(250);column:name"`
	Description string `json:"description,omitempty" gorm:"type:varchar(5000);column:description"`
	Status      string `json:"status,omitempty" gorm:"type:varchar(20);column:status"`
	CreatedBy   string `json:"createdBy,omitempty" gorm:"type:varchar(50);column:created_by"`
	UpdatedBy   string `json:"updatedBy,omitempty" gorm:"type:varchar(50);colum:updated_by"`
}

func (ProductType) TableName() string {
	return "product_type"
}
