package db

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
)

func MySql() *gorm.DB {
	loadEnvError := godotenv.Load()

	if loadEnvError != nil {
		panic("Can't read .env file")
	}

	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DATABASE")
	userName := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")

	hostConnection := userName + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8&parseTime=true&loc=UTC"

	db, openDbError := gorm.Open(mysql.Open(hostConnection), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase: true,
		},
	})

	if openDbError != nil {
		panic("Can't connect to database")
	}

	fmt.Println("Connected using MySQL")

	return db
}
