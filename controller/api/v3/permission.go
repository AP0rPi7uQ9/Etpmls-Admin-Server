package v3

import (
	"Etpmls-Admin-Server/core"
	"Etpmls-Admin-Server/library"
	"Etpmls-Admin-Server/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Create Permission
// 创建允许
func PermissionCreate(c *gin.Context)  {
	var j model.ApiPermissionCreate

	// Bind data
	if err := c.ShouldBindJSON(&j); err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_PermissionCreate_Bind, core.ERROR_MESSAGE_PermissionCreate_Bind, nil, err)
		return
	}

	// Validate Form
	err := library.ValidateZh(j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_PermissionCreate_Validate, core.ERROR_MESSAGE_PermissionCreate_Validate, nil, err)
		return
	}

	var p model.Permission
	err = p.PermissionCreate(j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_PermissionCreate, core.ERROR_MESSAGE_PermissionCreate, nil, err)
		return
	}

	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_PermissionCreate, core.SUCCESS_MESSAGE_PermissionCreate, nil)
	return
}

// Get all Permission
// 获取所有的权限
func PermissionGetAll(c *gin.Context)  {
	var p model.Permission
	data, count := p.PermissionGetAll(c)

	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_PermissionGetAll, core.SUCCESS_MESSAGE_PermissionGetAll, gin.H{"data": data, library.Config.App.Api.Pagination.Field.Count: count})
	return
}

// Modify Permission
// 修改权限
func PermissionEdit(c *gin.Context)  {
	var j model.ApiPermissionEdit

	// Bind data
	if err := c.ShouldBindJSON(&j); err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_PermissionEdit_Bind, core.ERROR_MESSAGE_PermissionEdit_Bind, nil, err)
		return
	}

	// Validate Form
	err := library.ValidateZh(j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_PermissionEdit_Validate, core.ERROR_MESSAGE_PermissionEdit_Validate, nil, err)
		return
	}

	var p model.Permission
	err = p.PermissionEdit(j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_PermissionEdit, core.ERROR_MESSAGE_PermissionEdit, nil, err)
		return
	}

	// Return Message
	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_PermissionEdit, core.SUCCESS_MESSAGE_PermissionEdit, nil)
	return
}

// Delete Permission (multiple can be deleted at the same time)
// 删除权限(可以同时删除多个)
func PermissionDelete(c *gin.Context)  {
	var j model.ApiPermissionDelete
	// Bind data
	if err := c.ShouldBindJSON(&j); err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_PermissionDelete_Bind, core.ERROR_MESSAGE_PermissionDelete_Bind, nil, err)
		return
	}

	var ids []uint
	for _, v := range j.Permissions {
		ids = append(ids, v.ID)
	}

	var p model.Permission
	err := p.PermissionDelete(ids)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_PermissionDelete, core.ERROR_MESSAGE_PermissionDelete, nil, err)
		return
	}

	// Return Message
	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_PermissionDelete, core.SUCCESS_MESSAGE_PermissionDelete, nil)
	return
}