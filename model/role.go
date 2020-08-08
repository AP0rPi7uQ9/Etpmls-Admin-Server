package model

import (
	"Etpmls-Admin-Server/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type Role struct {
	ID        uint `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Name string `json:"name"`
	Remark string `json:"remark"`
	Users []User `gorm:"many2many:role_users" json:"users"`
	Permissions []Permission `gorm:"many2many:role_permissions;" json:"permissions"`
}

type ApiRoleCreateV2 struct {
	ID        uint `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Name string `json:"name" binding:"required" validate:"max=30"`
	Remark string `json:"remark"`
	Permissions []Permission `gorm:"many2many:role_permissions;" json:"permissions"`
}

type ApiRoleEditV2 struct {
	ID        uint `json:"id" binding:"required" validate:"min=1"`
	CreatedAt time.Time `gorm:"-" json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Name string `json:"name" binding:"required" validate:"max=30"`
	Remark string `json:"remark"`
	Permissions []Permission `gorm:"many2many:role_permissions;" json:"permissions"`
}

type ApiRoleGetAllV2 struct {
	ID        uint `json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Name string `json:"name" binding:"required" validate:"max=30"`
	Remark string `json:"remark"`
	Permissions []Permission `gorm:"many2many:role_permissions;" json:"permissions"`
}

type ApiRoleDeleteV2 struct {
	ID uint `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
	Roles []Role `json:"roles" binding:"required" validate:"min=1"`
}

// Create Role
// 创建角色
func RoleCreateV2(j ApiRoleCreateV2) (error) {
	type Role ApiRoleCreateV2
	form := Role(j)

	// Insert Data
	result := database.DB.Create(&form)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Modify role
// 修改角色
func RoleEditV2(j ApiRoleEditV2) (error) {
	type Role ApiRoleEditV2
	form := Role(j)
	result := database.DB.Save(&form)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func RoleGetAllV2(c *gin.Context) (interface{}, int64) {
	type Role ApiRoleGetAllV2
	var data []Role

	limit, offset := CommonGetPageByQuery(c)
	var count int64
	// Get the title of the search, if not get all the data
	// 获取搜索的标题，如果没有获取全部数据
	search := c.Query("search")

	database.DB.Model(&Role{}).Preload("Permissions").Where("name " + database.FUZZY_SEARCH + " ?", "%"+ search +"%").Count(&count).Limit(limit).Offset(offset).Find(&data)

	return data, count
}

// Delete roles (allow multiple deletions at the same time)
// 删除角色（允许同时删除多个）
func RoleDeleteV2(ids []uint) (err error) {
	err = database.DB.Transaction(func(tx *gorm.DB) error {
		var r []Role
		database.DB.Where("id IN ?", ids).Find(&r)

		// 删除角色
		result := database.DB.Debug().Where("id IN ?", ids).Delete(&Role{})
		if result.Error != nil {
			return result.Error
		}

		// 删除关联
		err = database.DB.Model(&r).Association("Users").Clear()
		if err != nil {
			return err
		}

		// 删除关联
		err = database.DB.Model(&r).Association("Permissions").Clear()
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
