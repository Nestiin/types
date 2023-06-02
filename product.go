package types

import (
	"github.com/jinzhu/gorm"
)

// Product
type Product struct {
	gorm.Model
	Code          string               `json:"code,omitempty" gorm:"type:varchar(250);column:code"`
	Name          string               `json:"name,omitempty" gorm:"type:varchar(250);column:name"`
	BranchID      int                  `json:"branchId,omitempty" gorm:"type:int;column:branch_id"`
	CategoryID    uint                 `json:"categoryId,omitempty" gorm:"type:int;column:category_id"`
	TypeID        uint                 `json:"typeId,omitempty" gorm:"type:int;column:type_id"`
	SubCategoryID uint                 `json:"subCategoryId,omitempty" gorm:"type:int;column:subcategory_id"`
	Description   string               `json:"description,omitempty" gorm:"type:varchar(1000);column:description"`
	Thumbnail     Image                `json:"thumbnail,omitempty" gorm:"-"`
	Images        []Image              `json:"images,omitempty" gorm:"-"`
	IsFragile     bool                 `json:"isFragile,omitempty" gorm:"type:bool;column:is_fragile"`
	UOMID         int                  `json:"uomID,omitempty" gorm:"type:int;column:UOM_id"`
	Gender        string               `json:"gender,omitempty" gorm:"type:varchar(100);column:gender"`
	AgeGroup      string               `json:"ageGroup,omitempty" gorm:"type:varchar(50);column:age_group"`
	MerchantCode  string               `json:"merchantCode,omitempty" gorm:"type:varchar(50);column:merchant_code"`
	Variants      []ProductVariantInfo `json:"variants,omitempty" validate:"dive" gorm:"ForeignKey:ProductId"`
	MediaDetails  []string             `json:"mediaDetails,omitempty" gorm:"-"`
	MediaInfo     []MediaInfo          `json:"mediaInfo,omitempty" gorm:"ForeignKey:ProductId"`
	Length        int                  `json:"length,omitempty" gorm:"type:int;column:length"`
	Status        string               `json:"status,omitempty" gorm:"type:varchar(50);column:status"`
	Comment       string               `json:"comment,omitempty" gorm:"type:varchar(250);column:review_comment"`
}

func (Product) TableName() string {
	return "product_info"
}

type Image struct {
	Caption string `json:"caption"`
	URL     string `json:"url"`
	Height  string `json:"height"`
	Width   string `json:"width"`
}
