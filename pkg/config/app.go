// Returns db variable, helps other files to interact with db
package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var (
	db *gorm.DB
)

func Connect() {

	err := godotenv.Load()

	if err != nil {
		fmt.Println(os.Getwd())
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	// dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")

	//gorm.Open(dbDriver,dbUser:dbPassword/database_name?..)

	d, err := gorm.Open("mysql", dbUser+":"+dbPassword+"@tcp(localhost:3306)/simplerest?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d

}

func GetDB() *gorm.DB {
	return db
}
