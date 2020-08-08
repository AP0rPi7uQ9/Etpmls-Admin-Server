package main

import (
	_ "Etpmls-Admin-Server/database"
	"Etpmls-Admin-Server/library"
	"Etpmls-Admin-Server/middleware"
	"Etpmls-Admin-Server/route"
	"Etpmls-Admin-Server/utils/initialization"
	"github.com/gin-gonic/gin"
)

func main() {
	library.InitLogrus()
	library.InitRedis()
	initialization.InitDatabase()
	router := initRoute()
	_ = router.Run(":" + library.Config.App.Port)
}

func initRoute() *gin.Engine {
	router := gin.Default()

	// Load Front End Files
	initRouteStatic(router)

	// WEB Route
	route.RouteWeb(router)

	// Middleware - CORS
	middleware.InitCors(router)

	// API Route
	route.RouteApi(router)

	return router
}

func initRouteStatic(router *gin.Engine) {
	router.LoadHTMLGlob("storage/view/dist/*.html")
	router.Static("/static", "storage/view/static")
	router.Static("/storage/upload", "storage/upload")
	return
}
