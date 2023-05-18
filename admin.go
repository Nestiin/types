package types

import "github.com/jinzhu/gorm"

type AdminRegister struct {
	gorm.Model
	AdminCode        string ` gorm:"type:varchar(50);column:admin_code"`
	FirstName        string `json:"first_name" gorm:"type:varchar(45);column:first_name"`
	LastName         string `json:"last_name" gorm:"type:varchar(45);column:last_name"`
	RoleID           uint   `json:"role_id" gorm:"type:int;column:role_id"`
	UserName         string `json:"user_name" gorm:"type:varchar(45);column:user_name"`
	PhoneNo          string `json:"phone" gorm:"type:varchar(20);column:phone"`
	Email            string `json:"email" gorm:"type:varchar(45);column:email"`
	Password         string `json:"password" gorm:"type:varchar(45);column:password"`
	IsFirstTimeLogin bool   `json:"is_first_time_login" gorm:"type:boolean;column:is_first_time_login"`
	IsTokenVerified  bool   `json:"is_token_verified" gorm:"type:boolean;column:is_token_verified"`
	IsActive         bool   `json:"is_active" gorm:"type:varchar;column:is_active"`
	CreatedBy        string `json:"created_by" gorm:"type:varchar(50);column:created_by"`
	UpdatedBy        string `json:"updated_by" gorm:"type:varchar(50);column:updated_by"`
}

func (AdminRegister) TableName() string {
	return "admin_info"
}

type AdminLoginReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
