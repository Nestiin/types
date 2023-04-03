package types

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Merchant
type Merchant struct {
	gorm.Model
	MerchantName       string `json:"name" gorm:"type:varchar(250);column:merchant_name"`
	MerchantCode       string `json:"code" gorm:"type:varchar(50);column:merchant_code"`
	ContactName        string `json:"contact" gorm:"type:varchar(15);column:contact_name"`
	PhoneNo            string `json:"phone" gorm:"type:varchar(50);column:phone_no"`
	Email              string `json:"email" gorm:"type:varchar(255);column:email"`
	Password           string `json:"password" gorm:"type:varchar(50);column:password"`
	IsTokenVerified    bool   `json:"isTokenVerified" gorm:"type:boolean;column:is_token_verified"`
	Isactive           bool   `json:"isActive" gorm:"type:varchar;column:isactive"`
	PANNumber          string `json:"pan" gorm:"type:varchar(20);column:pan_number"`
	GSTNumber          string `json:"gst" gorm:"type:varchar(20);column:gst_number"`
	RegistrationNumber string `json:"registrationNo" gorm:"type:varchar(20);column:registration_number"`
	IsAdminVerified    bool   `json:"isAdminVerified"  gorm:"column:is_admin_verified"`
	Role               string `json:"role" gorm:"type:varchar(10);column:role"`
	Token              string `json:"token" gorm:"type:varchar(20);column:token"`
	Reason             string `json:"reason" gorm:"type:varchar(250);column:reason"`
	CreatedBy          string `json:"createdBy" gorm:"type:varchar(50);column:created_by"`
	UpdatedBy          string `json:"updatedBy" gorm:"type:varchar(50);column:updated_by"`
}

func (Merchant) TableName() string {
	return "merchant_info"
}


func SayHello(){
    fmt.Println("Hello From package")
}


