package core

import (
	"Etpmls-Admin-Server/library"
	"github.com/gin-gonic/gin"
)

// Error Code
// 错误码
const (
	ERROR_Code = "400000"
)

// Error Message
// 错误信息
const (
	ERROR_MESSAGE_Login = "登录失败！"
	ERROR_MESSAGE_Register = "注册失败！"
	ERROR_MESSAGE_RegistrationClosed = "管理员关闭了注册功能！"
	ERROR_MESSAGE_Logout = "登出失败！"
	ERROR_MESSAGE_Create = "创建失败！"
	ERROR_MESSAGE_Edit = "更新失败！"
	ERROR_MESSAGE_Get = "获取失败！"
	ERROR_MESSAGE_Delete = "删除失败！"
	ERROR_MESSAGE_Upload = "上传失败！"
	ERROR_MESSAGE_BindData = "提交的参数存在错误"
	ERROR_MESSAGE_Validate = "验证失败"
	ERROR_MESSAGE_PermissionDenied = "权限不足"
	ERROR_MESSAGE_DuplicateUserName = "当前用户名已存在！"
	ERROR_MESSAGE_DuplicateRoleName = "当前角色名已存在！"
	ERROR_MESSAGE_ProhibitOperationOfAdministratorUsers = "禁止操作管理员用户！"
	ERROR_MESSAGE_GetToken = "获取令牌失败！"
	ERROR_MESSAGE_TokenVerificationFailed = "令牌校验失败！"
)

// Success Code
// 成功码
const (
	SUCCESS_Code = "200000"
)

// Success Message
// 成功信息
const (
	SUCCESS_MESSAGE_Login = "登录成功！"
	SUCCESS_MESSAGE_Register = "注册成功！"
	SUCCESS_MESSAGE_Logout = "登出成功！"
	SUCCESS_MESSAGE_Create = "创建成功！"
	SUCCESS_MESSAGE_Edit = "更新成功！"
	SUCCESS_MESSAGE_Get = "获取成功！"
	SUCCESS_MESSAGE_Delete = "删除成功！"
	SUCCESS_MESSAGE_Upload = "上传成功！"
)

var (
	api_code = library.Config.App.Api.Field.Code
	api_status = library.Config.App.Api.Field.Status
	api_message = library.Config.App.Api.Field.Message
	api_data = library.Config.App.Api.Field.Data
)


// Return error message in json format
// 返回json格式的错误信息
func JsonError(c *gin.Context, httpStatusCode int, code interface{}, message string, data interface{}, err error)  {
	// If enabled, use HTTP CODE instead of system default CODE
	// 如果开启使用HTTP CODE 代替系统的默认CODE
	if library.Config.App.Api.UseHttpCode == true {
		code = httpStatusCode
	}

	// If it is a Debug environment, return information with Error
	// 如果是Debug环境，返回带有Error的信息
	if IsDebug() && err != nil {
		c.JSON(httpStatusCode, gin.H{api_code: code, api_status:"error", api_message: message + "Error: " + err.Error(), api_data: data})
		return
	}

	c.JSON(httpStatusCode, gin.H{api_code: code, api_status:"error", api_message: message, api_data: data})
	return
}


// Return success information in json format
// 返回json格式的成功信息
func JsonSuccess(c *gin.Context, httpStatusCode int, code interface{}, message string, data interface{})  {
	// If enabled, use HTTP CODE instead of system default CODE
	// 如果开启使用HTTP CODE 代替系统的默认CODE
	if library.Config.App.Api.UseHttpCode == true {
		code = httpStatusCode
	}

	c.JSON(httpStatusCode, gin.H{api_code: code, api_status:"success", api_message: message, api_data: data})
	return
}