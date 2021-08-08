package database

import (
	"fmt"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

const DB_USERNAME = "root"
const DB_PASSWORD = "test"
const DB_NAME = "Godb"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	var err error
	dsn := "root:test@tcp(127.0.0.1:3306)/GODB?parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Eror connecting to database : error=%v", err)
		return nil
	}

	return db
}
