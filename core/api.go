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


// Success Code
// 成功码
const (
	SUCCESS_Code = "200000"
)


var (
	api_code = library.Config.Field.Api.Code
	api_status = library.Config.Field.Api.Status
	api_message = library.Config.Field.Api.Message
	api_data = library.Config.Field.Api.Data
)


// Return error message in json format
// 返回json格式的错误信息
func JsonError(c *gin.Context, httpStatusCode int, code interface{}, message string, data interface{}, err error)  {
	// If enabled, use HTTP CODE instead of system default CODE
	// 如果开启使用HTTP CODE 代替系统的默认CODE
	if library.Config.App.UseHttpCode == true {
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
	if library.Config.App.UseHttpCode == true {
		code = httpStatusCode
	}

	c.JSON(httpStatusCode, gin.H{api_code: code, api_status:"success", api_message: message, api_data: data})
	return
}