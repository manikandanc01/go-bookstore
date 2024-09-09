// Returns db variable, helps other files to interact with db
package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	//gorm.Open(dbDriver,dbUser:dbPassword/database_name?..)
	// "root:mysqlHeaven@1234@tcp(localhost:3306)/simplerest?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", "root:mysqlHeaven@1234@tcp(localhost:3306)/simplerest?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
	
}

func GetDB() *gorm.DB {
	return db
}
