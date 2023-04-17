package types

import "github.com/jinzhu/gorm"

type MediaInfo struct {
	gorm.Model
	ProductId int    `json:"productId" gorm:"type:bigint;column:product_id"`
	MediaId   string `json:"mediaId" gorm:"type:text;column:media_id"`
	IsActive  bool   `json:"isActive" gorm:"type:bool;column:isactive"`
}

func (MediaInfo) TableName() string {
	return "media_info"
}
