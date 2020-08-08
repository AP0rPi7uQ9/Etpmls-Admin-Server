package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	APP_Version = "1.0.0"
	APP_Name = "Etmpls-Admin"
)

func GetAppDetail(c *gin.Context)  {
	ReturnJsonSuccess(c, http.StatusOK, 0, nil, gin.H{"version": APP_Version, "name": APP_Name})
	return
}

func ReturnJsonError(c *gin.Context, httpStatusCode int, code int, message, data interface{})  {
	c.JSON(httpStatusCode, gin.H{"code": code, "status":"error", "message": message, "data": data})
	return
}

func ReturnJsonSuccess(c *gin.Context, httpStatusCode int, code int, message, data interface{})  {
	c.JSON(httpStatusCode, gin.H{"code": code, "status":"success", "message": message, "data": data})
	return
}
