// +build mysql

package database

import (
	"Etpmls-Admin1-Server/library"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

const (
	FUZZY_SEARCH = "LIKE"
)

func init() {
	var dsn string

	dsn =  library.Config.Database.User + ":" + library.Config.Database.Password + "@(" + library.Config.Database.Host + ":" + library.Config.Database.Port + ")/" + library.Config.Database.Name + "?charset=utf8&parseTime=True&loc=Local"


	//Connect Database
	var err error
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		library.Log.Error(err)
		panic("failed to connect database")
	}


	//Table Prefix
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return library.Config.Database.Prefix + defaultTableName;
	}

	// Add table suffix when create tables
	m := []interface{}{&User{}, &Menu{}, &Role{}, &Attachment{}}
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(m...)


}