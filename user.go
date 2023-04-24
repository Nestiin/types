package types

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	CountryCode     string `json:"countryCode" gorm:"type:varchar(50);column:country_code"`
	MemberCode      string `gorm:"type:varchar(50);column:member_code"`
	FirstName       string `json:"firstName" validate:"required" gorm:"type:varchar(50);column:first_name"`
	LastName        string `json:"lastName,omitempty" gorm:"type:varchar(50);column:last_name"`
	PhoneNo         string `json:"phoneNo" validate:"required,min=9,max=13,numeric" gorm:"type:varchar(15);column:phone_no"`
	TokenVerified   bool   `gorm:"type:boolean;column:is_token_verified"`
	Email           string `json:"email,omitempty" gorm:"type:varchar(255);column:email"`
	Password        string `json:"password" validate:"required,min=8,max=20" gorm:"type:varchar(50);column:password"`
	Address         string `json:"address" gorm:"type:varchar(2000);column:address"`
	DeliveryAddress string `json:"deliveryAddress" gorm:"type:varchar(2000);column:delivery_address"`
	SocialMediaID   string `json:"socialMediaId" gorm:"type:varchar(100);column:social_media_id"`
	SocialMediaType string `json:"socialMediaType" gorm:"type:varchar(100);column:social_media_type"`
	Status          string `gorm:"type:varchar(50);column:status"`
	RegisterBy      string `json:"registerBy" gorm:"-"`
}

func (User) TableName() string {
	return "user_info"
}
