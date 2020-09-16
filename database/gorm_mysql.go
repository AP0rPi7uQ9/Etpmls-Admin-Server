// +build mysql

package database

import (
	"Etpmls-Admin-Server/core"
	"Etpmls-Admin-Server/library"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

const (
	FUZZY_SEARCH = "LIKE"
)

func init() {
	dsn := library.Config.Database.User + ":" + library.Config.Database.Password + "@tcp(" + library.Config.Database.Host + ":" + library.Config.Database.Port + ")/" + library.Config.Database.Name + "?charset=utf8mb4&parseTime=True&loc=" +  + library.Config.App.TimeZone

	//Connect Database
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   library.Config.Database.Prefix,
		},
	})
	if err != nil {
		core.LogPanic.AutoOutputDebug("连接数据库失败！", err)
	}

	err = DB.AutoMigrate(migrateList...)
	if err != nil {
		core.LogInfo.AutoOutputDebug("创建数据库失败！", err)
	}

}