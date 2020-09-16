// +build postgresql

package database

import (
	"Etpmls-Admin-Server/core"
	"Etpmls-Admin-Server/library"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

const (
	FUZZY_SEARCH = "ILIKE"
)

func init() {
	dsn := "DB.host=" + library.Config.Database.Host + " user=" + library.Config.Database.User + " password=" + library.Config.Database.Password + " dbname=" + library.Config.Database.Name + " port=" + library.Config.Database.Port + " sslmode=disable TimeZone=" + library.Config.App.TimeZone

	//Connect Database
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
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