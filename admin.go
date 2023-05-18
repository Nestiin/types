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
	CreatePermission bool   `json:"create_permission" gorm:"-"`
	UpdatePermission bool   `json:"update_permission" gorm:"-"`
	ReadPermission   bool   `json:"read_permission" gorm:"-"`
	DeletePermission bool   `json:"delete_permission" gorm:"-"`
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

type MemberRole struct {
	gorm.Model
	MemberRoleCode   string `gorm:"type:varchar(50);column:member_role_code"`
	RoleID           uint   `gorm:"type:int;column:role_id"`
	MemberCode       string `gorm:"type:varchar(50);column:member_code"`
	IsActive         bool   `gorm:"type:varchar;column:is_active"`
	CreatePermission bool   `gorm:"type:varchar;column:create_permission"`
	UpdatePermission bool   `gorm:"type:varchar;column:update_permission"`
	ReadPermission   bool   `gorm:"type:varchar;column:read_permission"`
	DeletePermission bool   `gorm:"type:varchar;column:delete_permission"`
	CreatedBy        string `gorm:"type:varchar(50);column:created_by"`
	UpdatedBy        string `gorm:"type:varchar(50);column:updated_by"`
	Role             Role   `gorm:"references:id"`
}

func (MemberRole) TableName() string {
	return "member_role"
}

type Role struct {
	gorm.Model
	RoleCode    string `gorm:"type:varchar(50);column:role_code"`
	RoleName    string `gorm:"type:varchar(50);column:role_name"`
	IsActive    bool   `gorm:"type:varchar;column:is_active"`
	Permissions string `gorm:"type:text;column:permissions"`
	CreatedBy   string `gorm:"type:varchar(50);column:created_by"`
	UpdatedBy   string `gorm:"type:varchar(50);column:updated_by"`
}

func (Role) TableName() string {
	return "role"
}
