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
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_BindData"), nil, err)
		return
	}

	// Validate Form
	err := library.ValidateZh(j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_Validate"), nil, err)
		return
	}

	var p model.Permission
	err = p.PermissionCreate(j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_Create"), nil, err)
		return
	}

	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_Code, core.Translate(c, "SUCCESS_MESSAGE_Create"), nil)
	return
}

// Get all Permission
// 获取所有的权限
func PermissionGetAll(c *gin.Context)  {
	var p model.Permission
	data, count := p.PermissionGetAll(c)

	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_Code, core.Translate(c, "SUCCESS_MESSAGE_Get"), gin.H{"data": data, library.Config.Field.Pagination.Count: count})
	return
}

// Modify Permission
// 修改权限
func PermissionEdit(c *gin.Context)  {
	var j model.ApiPermissionEdit

	// Bind data
	if err := c.ShouldBindJSON(&j); err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_BindData"), nil, err)
		return
	}

	// Validate Form
	err := library.ValidateZh(j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_Validate"), nil, err)
		return
	}

	var p model.Permission
	err = p.PermissionEdit(j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_Edit"), nil, err)
		return
	}

	// Return Message
	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_Code, core.Translate(c, "SUCCESS_MESSAGE_Edit"), nil)
	return
}

// Delete Permission (multiple can be deleted at the same time)
// 删除权限(可以同时删除多个)
func PermissionDelete(c *gin.Context)  {
	var j model.ApiPermissionDelete
	// Bind data
	if err := c.ShouldBindJSON(&j); err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_BindData"), nil, err)
		return
	}

	var ids []uint
	for _, v := range j.Permissions {
		ids = append(ids, v.ID)
	}

	var p model.Permission
	err := p.PermissionDelete(ids)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_Delete"), nil, err)
		return
	}

	// Return Message
	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_Code, core.Translate(c, "SUCCESS_MESSAGE_Delete"), nil)
	return
}