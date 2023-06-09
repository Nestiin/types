package types

import "github.com/jinzhu/gorm"

type AdminRegister struct {
	gorm.Model
	Code             string `json:"code" gorm:"type:varchar(50);column:admin_code"`
	FirstName        string `json:"firstName" gorm:"type:varchar(45);column:first_name"`
	LastName         string `json:"lastName" gorm:"type:varchar(45);column:last_name"`
	RoleID           uint   `json:"roleId" gorm:"type:int;column:role_id"`
	UserName         string `json:"userName" gorm:"type:varchar(45);column:user_name"`
	PhoneNo          string `json:"phone" gorm:"type:varchar(20);column:phone"`
	Email            string `json:"email" gorm:"type:varchar(45);column:email"`
	Password         string `json:"password" gorm:"type:varchar(45);column:password"`
	IsFirstTimeLogin bool   `json:"isFirstTimeLogin" gorm:"type:boolean;column:is_first_time_login"`
	IsTokenVerified  bool   `json:"isTokenVerified" gorm:"type:boolean;column:is_token_verified"`
	Status           string `json:"status" gorm:"type:varchar;column:status"`
	CreatePermission bool   `json:"createPermission" gorm:"-"`
	UpdatePermission bool   `json:"updatePermission" gorm:"-"`
	ReadPermission   bool   `json:"readPermission" gorm:"-"`
	DeletePermission bool   `json:"deletePermission" gorm:"-"`
	CreatedBy        string `json:"createdBy" gorm:"type:varchar(50);column:created_by"`
	UpdatedBy        string `json:"updatedBy" gorm:"type:varchar(50);column:updated_by"`
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
	Status           string `gorm:"type:varchar;column:status"`
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
	RoleCode    string `json:"roleCode" gorm:"type:varchar(50);column:role_code"`
	RoleName    string `json:"roleName" gorm:"type:varchar(50);column:role_name"`
	Status      string `json:"status" gorm:"type:text;column:status"`
	Permissions string `json:"permissions" gorm:"type:text;column:permissions"`
	CreatedBy   string `json:"createdBy" gorm:"type:varchar(50);column:created_by"`
	UpdatedBy   string `json:"updatedBy" gorm:"type:varchar(50);column:updated_by"`
}

func (Role) TableName() string {
	return "role"
}

type UpdatePermission struct {
	RoleID           uint   `json:"roleId" gorm:"type:integer;column:role_id"`
	MemberRoleID     uint   `json:"id" validate:"required" gorm:"type:integer;column:member_role_id"`
	Status           string `json:"status" validate:"required" gorm:"type:varchar(50);column:status"`
	CreatePermission bool   `json:"createPermission" gorm:"type:bool;column:create_permission"`
	UpdatePermission bool   `json:"updatePermission" gorm:"type:bool;column:update_permission"`
	ReadPermission   bool   `json:"readPermission" gorm:"type:bool;column:read_permission"`
	DeletePermission bool   `json:"deletePermission" gorm:"type:bool;column:delete_permission"`
}

// Used to create new password
type CreateNewPassword struct {
	UserName        string `json:"username" validate:"required"`
	OldPassword     string `json:"oldPassword"`
	NewPassword     string `json:"newPassword" validate:"required"`
	ConfirmPassword string `json:"confirmPassword" validate:"required"`
}
