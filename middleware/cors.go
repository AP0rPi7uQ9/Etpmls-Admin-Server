package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitCors(router *gin.Engine)  {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "X-Token")
	config.AllowHeaders = append(config.AllowHeaders, "token")
	config.AllowHeaders = append(config.AllowHeaders, "Token")
	config.AllowHeaders = append(config.AllowHeaders, "language")
	router.Use(cors.New(config))
	return
}