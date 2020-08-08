package model

import (
	"Etpmls-Admin-Server/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Permission struct {
	ID        uint `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Name string `json:"name"`
	Method string `json:"method"`
	Path	string	`json:"path"`
	Remark string `json:"remark"`
	Roles []Role `gorm:"many2many:role_permissions" json:"roles"`
}

type ApiPermissionCreateV2 struct {
	ID        uint `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Name string `json:"name" binding:"required" validate:"max=255"`
	Method string `json:"-"`
	Path	string	`json:"path" binding:"required" validate:"max=255"`
	Remark string `json:"remark"`
	TmpMethod []string `gorm:"-" json:"method" binding:"required" validate:"min=1"`
}

type ApiPermissionGetAllV2 struct {
	ID        uint `json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Name string `json:"name"`
	Method string `json:"method"`
	Path	string	`json:"path"`
	Remark string `json:"remark"`
	Roles []Role `gorm:"many2many:role_permissions" json:"roles"`
}

type ApiPermissionEditV2 struct {
	ID        uint `json:"id" binding:"required" validate:"min=1"`
	CreatedAt time.Time `gorm:"-" json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Name string `json:"name" binding:"required" validate:"max=255"`
	Method string `json:"-"`
	Path	string	`json:"path" binding:"required" validate:"max=255"`
	Remark string `json:"remark"`
	TmpMethod []string `gorm:"-" json:"method" binding:"required" validate:"min=1"`
}

type ApiPermissionDeleteV2 struct {
	ID uint `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
	Permissions []Permission `json:"permissions" binding:"required" validate:"min=1"`
}

// Create Permission
// 创建权限
func PermissionCreateV2(j ApiPermissionCreateV2) (error) {
	type Permission ApiPermissionCreateV2
	form := Permission(j)

	// []string -> string
	form.Method = strings.Join(form.TmpMethod, ",")

	// Insert Data
	result := database.DB.Create(&form)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Get all Permission
// 获取所有的权限
func PermissionGetAllV2(c *gin.Context) (interface{}, int64) {
	type Permission ApiPermissionGetAllV2
	var data []Permission

	limit, offset := CommonGetPageByQuery(c)
	var count int64
	// Get the title of the search, if not get all the data
	// 获取搜索的标题，如果没有获取全部数据
	search := c.Query("search")

	database.DB.Model(&Permission{}).Where("name " + database.FUZZY_SEARCH + " ?", "%"+ search +"%").Count(&count).Limit(limit).Offset(offset).Find(&data)

	return data, count
}

// Modify Permission
// 修改权限
func PermissionEditV2(j ApiPermissionEditV2) (error) {
	type Permission ApiPermissionEditV2
	form := Permission(j)

	// []string -> string
	form.Method = strings.Join(form.TmpMethod, ",")

	result := database.DB.Save(&form)
	if result.Error != nil {
		return result.Error
	}

	return nil
}


// Delete Permission (allow multiple deletions at the same time)
// 删除权限（允许同时删除多个）
func PermissionDeleteV2(ids []uint) (err error) {
	err = database.DB.Transaction(func(tx *gorm.DB) error {
		var p []Permission
		database.DB.Where("id IN ?", ids).Find(&p)

		// 删除权限
		result := database.DB.Delete(&p)
		if result.Error != nil {
			return result.Error
		}

		// 删除关联
		err = database.DB.Model(&p).Association("Roles").Clear()
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
