// +build postgresql

package database

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Avatar Attachment	`gorm:"polymorphic:Owner;polymorphicValue:user-avatar"`
	Roles []Role `gorm:"many2many:role_users;"`
}

type Role struct {
	gorm.Model
	Name string
	Remark string
	Users []User `gorm:"many2many:role_users;"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}

type Permission struct {
	gorm.Model
	Name string
	Method string
	Path	string
	Remark string
	Roles []Role `gorm:"many2many:role_permissions;"`
}


/*type Menu struct {
	gorm.Model
	ParentId uint `gorm:"default:0;not null"`
	Hidden bool `gorm:"default:false;not null"`
	Redirect string
	AlwaysShow bool `gorm:"default:false;not null"`
	Name string
	Title string
	Icon string
	NoCache bool `gorm:"default:false;not null"`
	Affix bool `gorm:"default:false;not null"`
	Breadcrumb bool `gorm:"default:true;not null"`
	ActiveMenu string
	Path string
	Component string
	Roles []Role  `gorm:"many2many:role_menus"`
}*/

type Attachment struct {
	gorm.Model
	Path string	`gorm:"type:varchar(500)"`
	OwnerID uint
	OwnerType string
}