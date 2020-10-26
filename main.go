package main

import (
	"Etpmls-Admin-Server/core"
	_ "Etpmls-Admin-Server/database"
	"Etpmls-Admin-Server/hook"
	"Etpmls-Admin-Server/library"
	"Etpmls-Admin-Server/middleware"
	"Etpmls-Admin-Server/module"
	"Etpmls-Admin-Server/route"
	"Etpmls-Admin-Server/utils"
	"Etpmls-Admin-Server/utils/initialization"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
)

var errChan = make(chan error)

func main() {
	initialization.InitDatabase()
	module.InitModule()
	router := initRoute()
	go init_MonitorExit()
	go init_RunServer(router)
	v := <- errChan
	core.LogInfo.Output(utils.MessageWithLineNum(v.Error()))

	// Run Hook After Exit
	var h hook.Hook
	h.ExitApplication()
}

func initRoute() *gin.Engine {
	router := gin.Default()

	// Load Front End Files
	initRouteStatic(router)

	// WEB Route
	route.RouteWeb(router)
	module.Module_RouteWeb(router)

	// Middleware - CORS
	middleware.InitCors(router)

	// API Route
	route.RouteApi(router)
	module.Module_RouteApi(router)

	return router
}

func initRouteStatic(router *gin.Engine) {
	router.LoadHTMLGlob("storage/view/dist/*.html")
	router.Static("/static", "storage/view/static")
	router.Static("/storage/upload", "storage/upload")
	return
}

func init_MonitorExit()  {
	s := make(chan os.Signal)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	errChan <- fmt.Errorf("%s", <-s)
}

func init_RunServer(router *gin.Engine)  {
	_ = router.Run(":" + library.Config.App.Port)
}