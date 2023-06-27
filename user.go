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

// user address global struct model
type UserAddress struct {
	gorm.Model
	UserId       uint   `json:"userId,omitempty" gorm:"type:bigserial;column:user_id"`
	AddressName  string `json:"addressName,omitempty" gorm:"type:varchar(50);column:address_name"`
	BlockNo      string `json:"blockNo,omitempty" gorm:"type:varchar(50);column:block_number"`
	Street       string `json:"street,omitempty" gorm:"type:varchar(50);column:street"`
	ApartmentNo  string `json:"apartmentNo,omitempty" gorm:"type:varchar(50);column:apartment_number"`
	BuildingName string `json:"buildingName,omitempty" gorm:"type:varchar(50);column:building_name"`
	Town         string `json:"town,omitempty" gorm:"type:varchar(50);column:town"`
	PostalCode   string `json:"postalCode,omitempty" gorm:"type:varchar(15);column:post_code"`
	State        string `json:"state,omitempty" gorm:"type:varchar(50);column:state"`
	Country      string `json:"country,omitempty" gorm:"type:varchar(50);column:country"`
	Status       string `json:"status,omitempty" gorm:"type:varchar(50);column:status"`
	IsPrimary    bool   `json:"isPrimary,omitempty" gorm:"type:boolean;column:is_primary"`
}

func (UserAddress) TableName() string {
	return "user_address"
}

type TokenPhoneRequest struct {
	Username string `json:"username"`
	Otp      string `json:"otp" validate:"required,min=4,max=6,alphanum"`
	Method   string `json:"method"`
}

type UpdateUserLog struct {
	gorm.Model
	Action  string `gorm:"type:varchar(256);column:action"`
	User_id uint   `gorm:"type:varchar(50);column:user_id"`
}

func (UpdateUserLog) TableName() string {
	return "user_log"
}

type ForgotPasswordReq struct {
	UserName string `json:"username"`
}

type NewPasswordRequest struct {
	UserName        string `json:"username"`
	OTP             string `json:"otp"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}
