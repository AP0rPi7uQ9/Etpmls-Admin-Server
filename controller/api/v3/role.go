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
	var j model.ApiRoleCreate

	// Bind data
	if err := c.ShouldBindJSON(&j); err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.ERROR_MESSAGE_BindData, nil, err)
		return
	}

	// Validate Form
	err := library.ValidateZh(j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.ERROR_MESSAGE_Validate, nil, err)
		return
	}

	// Validate Name Unique
	var count int64
	database.DB.Model(&model.Role{}).Where("name = ?", j.Name).Count(&count)
	if count > 0 {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.ERROR_MESSAGE_DuplicateRoleName, nil, err)
		return
	}

	var r model.Role
	err = r.RoleCreate(j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.ERROR_MESSAGE_Create, nil, err)
		return
	}

	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_Code, core.SUCCESS_MESSAGE_Create, nil)
	return
}

// Get all characters
// 获取所有的角色
func RoleGetAll(c *gin.Context)  {
	var r model.Role
	data, count := r.RoleGetAll(c)

	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_Code, core.SUCCESS_MESSAGE_Get, gin.H{"data": data, library.Config.App.Api.Pagination.Field.Count: count})
	return
}

// Modify role
// 修改角色
func RoleEdit(c *gin.Context)  {
	var j model.ApiRoleEdit

	// Bind data
	if err := c.ShouldBindJSON(&j); err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.ERROR_MESSAGE_BindData, nil, err)
		return
	}

	// Validate Form
	err := library.ValidateZh(j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.ERROR_MESSAGE_Validate, nil, err)
		return
	}

	// Validate Name Unique
	var count int64
	database.DB.Model(&model.Role{}).Where("name = ?", j.Name).Where("id != ?", j.ID).Count(&count)
	if count > 0 {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.ERROR_MESSAGE_DuplicateRoleName, nil, err)
		return
	}

	var r model.Role
	err = r.RoleEdit(j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.ERROR_MESSAGE_Edit, nil, err)
		return
	}

	// Return Message
	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_Code, core.SUCCESS_MESSAGE_Edit, nil)
	return
}

// Delete roles (multiple can be deleted at the same time)
// 删除角色(可以同时删除多个)
func RoleDelete(c *gin.Context)  {
	var j model.ApiRoleDelete
	// Bind data
	if err := c.ShouldBindJSON(&j); err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.ERROR_MESSAGE_BindData, nil, err)
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
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.ERROR_MESSAGE_ProhibitOperationOfAdministratorUsers, nil, nil)
		return
	}

	var r model.Role
	err := r.RoleDelete(ids)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.ERROR_MESSAGE_Delete, nil, err)
		return
	}

	// Return Message
	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_Code, core.SUCCESS_MESSAGE_Delete, nil)
	return
}
