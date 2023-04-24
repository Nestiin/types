package types

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Merchant
type Merchant struct {
	gorm.Model
	Name              string `json:"name,omitempty" gorm:"type:varchar(250);column:merchant_name"`
	Code              string `json:"code,omitempty" gorm:"type:varchar(50);column:merchant_code"`
	DisplayName       string `json:"displayName,omitempty" gorm:"type:varchar(250);column:display_name"`
	ContactName       string `json:"contactName,omitempty" gorm:"type:varchar(250);column:contact_name"`
	Email             string `json:"email,omitempty" gorm:"type:varchar(255);column:email"`
	HouseNumber       string `json:"houseNumber,omitempty" gorm:"type:varchar(13);column:house_number"`
	BuildingName      string `json:"buildingName,omitempty" gorm:"type:varchar(13);column:building_name"`
	StreetName        string `json:"streetName,omitempty" gorm:"type:VARCHAR(250);column:street_name"`
	City              string `json:"city,omitempty" gorm:"type:VARCHAR(250);column:city"`
	State             string `json:"state,omitempty" gorm:"type:VARCHAR(250);column:state"`
	Country           string `json:"country,omitempty" gorm:"type:VARCHAR(250);column:country"`
	PostalCode        string `json:"postalCode,omitempty" gorm:"type:VARCHAR(250);column:postal_code"`
	PANNumber         string `json:"panNumber,omitempty" gorm:"type:varchar(20);column:pan_number"`
	GSTNumber         string `json:"gstNumber,omitempty" gorm:"type:varchar(20);column:gst_number"`
	Password          string `json:"password,omitempty" gorm:"type:varchar(50);column:password"`
	PhoneNo           string `json:"phoneNo,omitempty" gorm:"type:varchar(13);column:phone_no"`
	IsAdminVerified   bool   `json:"isAdminVerified,omitempty" gorm:"column:is_admin_verified"`
	Status            string `json:"status,omitempty" gorm:"type:varchar(25);column:status"`
	RegistrationStage int    `json:"registrationStage,omitempty" gorm:"type:int;column:registration_stage"`
	IsTokenVerified   bool   `json:"isTokenVerified,omitempty" gorm:"type:boolean;column:is_token_verified"`
	Role              string `json:"role,omitempty" gorm:"type:varchar(10);column:role"`
	Token             string `json:"token,omitempty" gorm:"type:varchar(20);column:token"`
	Reason            string `json:"reason,omitempty" gorm:"type:varchar(250);column:reason"`
	RegisterBy        string `json:"registerBy,omitempty" gorm:"-"`
}

func (Merchant) TableName() string {
	return "merchant_info"
}

// Login request
type LoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func SayHello() {
	fmt.Println("Hello From package")
}
