package v3

import (
	"Etpmls-Admin-Server/core"
	"Etpmls-Admin-Server/database"
	"Etpmls-Admin-Server/library"
	"Etpmls-Admin-Server/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Create Role
// 创建角色
func RoleCreate(c *gin.Context)  {
	var j model.ApiRoleCreateV2

	// Bind data
	if err := c.ShouldBindJSON(&j); err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_RoleCreate_BIND, core.ERROR_MESSAGE_RoleCreate_BIND, nil, err)
		return
	}

	// Validate Form
	err := library.ValidateZh(j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_RoleCreate_VALIDATE, core.ERROR_MESSAGE_RoleCreate_VALIDATE, nil, err)
		return
	}

	// Validate Name Unique
	var count int64
	database.DB.Model(&model.Role{}).Where("name = ?", j.Name).Count(&count)
	if count > 0 {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_RoleCreate_ROLE_NAME_EXISTS, core.ERROR_MESSAGE_RoleCreate_ROLE_NAME_EXISTS, nil, err)
		return
	}

	err = model.RoleCreateV2(j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_RoleCreate_ROLE_CREATE, core.ERROR_MESSAGE_RoleCreate_ROLE_CREATE, nil, err)
		return
	}

	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_RoleCreate, core.SUCCESS_MESSAGE_RoleCreate, nil)
	return
}

// Get all characters
// 获取所有的角色
func RoleGetAll(c *gin.Context)  {
	data, count := model.RoleGetAllV2(c)

	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_RoleGetAll, core.SUCCESS_MESSAGE_RoleGetAll, gin.H{"data": data, library.Config.App.Api.Pagination.Field.Count: count})
	return
}

// Modify role
// 修改角色
func RoleEdit(c *gin.Context)  {
	var j model.ApiRoleEditV2

	// Bind data
	if err := c.ShouldBindJSON(&j); err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_RoleEdit_BIND, core.ERROR_MESSAGE_RoleEdit_BIND, nil, err)
		return
	}

	// Validate Form
	err := library.ValidateZh(j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_RoleEdit_VALIDATE, core.ERROR_MESSAGE_RoleEdit_VALIDATE, nil, err)
		return
	}

	// Validate Name Unique
	var count int64
	database.DB.Model(&model.Role{}).Where("name = ?", j.Name).Where("id != ?", j.ID).Count(&count)
	if count > 0 {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_RoleEdit_Duplicate_role, core.ERROR_MESSAGE_RoleEdit_Duplicate_role, nil, err)
		return
	}

	err = model.RoleEditV2(j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_RoleEdit, core.ERROR_MESSAGE_RoleEdit, nil, err)
		return
	}

	// Return Message
	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_RoleEdit, core.SUCCESS_MESSAGE_RoleEdit, nil)
	return
}

// Delete roles (multiple can be deleted at the same time)
// 删除角色(可以同时删除多个)
func RoleDelete(c *gin.Context)  {
	var j model.ApiRoleDeleteV2
	// Bind data
	if err := c.ShouldBindJSON(&j); err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_RoleDelete_Bind, core.ERROR_MESSAGE_RoleDelete_Bind, nil, err)
		return
	}

	var ids []uint
	for _, v := range j.Roles {
		ids = append(ids, v.ID)
	}
	fmt.Println(j.Roles)
	fmt.Println(ids)
	// Find if admin is included in ids
	// 查找ids中是否包含admin
	b := model.Common_CheckIfOneIsIncludeInIds(ids)
	if b {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_RoleDelete_Remove_admin_role, core.ERROR_MESSAGE_RoleDelete_Remove_admin_role, nil, nil)
		return
	}

	err := model.RoleDeleteV2(ids)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_RoleDelete, core.ERROR_MESSAGE_RoleDelete, nil, err)
		return
	}

	// Return Message
	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_RoleDelete, core.SUCCESS_MESSAGE_RoleDelete, nil)
	return
}
