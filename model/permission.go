package model

import (
	"Etpmls-Admin-Server/core"
	"Etpmls-Admin-Server/database"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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


// Create Permission
// 创建权限
type ApiPermissionCreate struct {
	ID        uint `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Name string `json:"name" binding:"required" validate:"max=255"`
	Method string `json:"method" binding:"required" validate:"min=1"`
	Path	string	`json:"path" binding:"required" validate:"max=255"`
	Remark string `json:"remark"`
}
func (this *Permission)PermissionCreate(c *gin.Context, j ApiPermissionCreate) (error) {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		type Permission ApiPermissionCreate
		form := Permission(j)

		// Insert Data
		result := tx.Create(&form)
		if result.Error != nil {
			core.LogError.Output(core.MessageWithLineNum(result.Error.Error()))
			return result.Error
		}

		// Create Event for module
		p, err := this.Permission_InterfaceToPermission(form)
		if err != nil {
			core.LogError.Output(core.MessageWithLineNum(err.Error()))
			return err
		}
		select {
		case core.Event.Event_PermissionCreate <- core.EventObject{
			Context: c,
			Content: p,
		}:
		case <- time.After(time.Second * 3):
		}


		return nil
	})


	return err
}


// Get all Permission
// 获取所有的权限
type ApiPermissionGetAll struct {
	ID        uint `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Name string `json:"name"`
	Method string `json:"method"`
	Path	string	`json:"path"`
	Remark string `json:"remark"`
	Roles []Role `gorm:"many2many:role_permissions" json:"roles"`
}
func (this *Permission) PermissionGetAll(c *gin.Context) (interface{}, int64) {
	type Permission ApiPermissionGetAll
	var data []Permission

	limit, offset := Common_GetPageByQuery(c)
	var count int64
	// Get the title of the search, if not get all the data
	// 获取搜索的标题，如果没有获取全部数据
	search := c.Query("search")

	database.DB.Model(&Permission{}).Where("name " + database.FUZZY_SEARCH + " ?", "%"+ search +"%").Count(&count).Limit(limit).Offset(offset).Find(&data)

	return data, count
}


// Modify Permission
// 修改权限
type ApiPermissionEdit struct {
	ID        uint `json:"id" binding:"required" validate:"min=1"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Name string `json:"name" binding:"required" validate:"max=255"`
	Method string `json:"method" binding:"required" validate:"min=1"`
	Path	string	`json:"path" binding:"required" validate:"max=255"`
	Remark string `json:"remark"`
}
func (this *Permission) PermissionEdit(c *gin.Context, j ApiPermissionEdit) (error) {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		type Permission ApiPermissionEdit
		form := Permission(j)


		result := tx.Save(&form)
		if result.Error != nil {
			core.LogError.Output(core.MessageWithLineNum(result.Error.Error()))
			return result.Error
		}

		// Edit Event for module
		p, err := this.Permission_InterfaceToPermission(form)
		if err != nil {
			return err
		}
		select {
		case core.Event.Event_PermissionEdit <- core.EventObject{
			Context: c,
			Content: p,
		}:
		case <- time.After(time.Second * 3):
		}

		return nil
	})


	return err
}


// Delete Permission (allow multiple deletions at the same time)
// 删除权限（允许同时删除多个）
type ApiPermissionDelete struct {
	ID uint `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
	Permissions []Permission `json:"permissions" binding:"required" validate:"min=1"`
}
func (this *Permission) PermissionDelete(c *gin.Context, ids []uint) (err error) {
	err = database.DB.Transaction(func(tx *gorm.DB) error {
		var p []Permission
		tx.Where("id IN ?", ids).Find(&p)

		// 删除权限
		result := tx.Delete(&p)
		if result.Error != nil {
			return result.Error
		}

		// 删除关联
		err = tx.Model(&p).Association("Roles").Clear()
		if err != nil {
			return err
		}

		// Delete Event for module
		select {
		case core.Event.Event_PermissionDelete <- core.EventObject{
			Context: c,
			Content: p,
		}:
		case <- time.After(time.Second * 3):
		}

		return nil
	})

	return err
}


// interface conversion Permission
// interface转换Permission
func (this *Permission) Permission_InterfaceToPermission(i interface{}) (Permission, error) {
	var p Permission
	us, err := json.Marshal(i)
	if err != nil {
		core.LogError.Output(core.MessageWithLineNum("Object to JSON failed!" + err.Error()))
		return Permission{}, err
	}
	err = json.Unmarshal(us, &p)
	if err != nil {
		core.LogError.Output(core.MessageWithLineNum("JSON conversion object failed!" + err.Error()))
		return Permission{}, err
	}
	return p, nil
}