package module

import (
	"Etpmls-Admin-Server/core"
	"Etpmls-Admin-Server/database"
)

var migrateList = []interface{}{}

func initDatabase()  {
	err := database.DB.AutoMigrate(migrateList...)
	if err != nil {
		core.LogInfo.AutoOutputDebug("创建数据库失败！", err)
	}
}