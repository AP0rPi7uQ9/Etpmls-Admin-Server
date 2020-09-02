package module

import (
	"Etpmls-Admin-Server/core"
	"Etpmls-Admin-Server/database"
)

// Registration database
// 注册数据库
// var migrateList = []interface{}{&cms_database.Category{}, &cms_database.Post{}, &cms_database.Variable{}}
var migrateList = []interface{}{}

func initDatabase()  {
	err := database.DB.AutoMigrate(migrateList...)
	if err != nil {
		core.LogInfo.AutoOutputDebug("创建数据库失败！", err)
	}
}

// Insert initial data
// 插入初始数据
func InsertModuleDataToDatabase()  {
	// cms_initialization.InsertCmsDataToDatabase()
}