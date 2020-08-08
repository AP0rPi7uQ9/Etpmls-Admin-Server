// +build postgresql-v1

package database

import (
	"Etpmls-Admin1-Server/library"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

const (
	FUZZY_SEARCH = "ILIKE"
)

func init() {
	var dsn string

	dsn = "host=" + library.Config.Database.Host + " port=" + library.Config.Database.Port + " user=" + library.Config.Database.User + " dbname=" + library.Config.Database.Name + " password=" + library.Config.Database.Password + " sslmode=disable"


	//Connect Database
	var err error
	DB, err = gorm.Open("postgres", dsn)
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
	DB.AutoMigrate(m...)


}