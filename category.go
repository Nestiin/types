package types

import "github.com/jinzhu/gorm"

type ProductCategory struct {
	gorm.Model
	Code        string `json:"code,omitempty" gorm:"type:varchar(50);column:code"`
	Name        string `json:"name,omitempty" gorm:"type:varchar(250);column:name"`
	Description string `json:"description,omitempty" gorm:"type:varchar(5000);column:description"`
	TypeID      int    `json:"typeID,omitempty" gorm:"type:int;column:type_id"`
	Status      string `json:"status,omitempty" gorm:"type:varchar(15);column:status"`
	CreatedBy   string `json:"createdBy,omitempty" gorm:"type:integer;column:created_by"`
}

func (ProductCategory) TableName() string {
	return "product_category"
}

type ProductCategoryAndSubCategory struct {
	ProductCategory
	Subcategories []ProductSubCategory `gorm:"foreignKey:category_id"`
}
