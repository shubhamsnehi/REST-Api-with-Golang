package database

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	// "gorm.io/driver/mysql"
	
)

var DB *gorm.DB

func Open() error {
	//dsn := "root:@tcp(localhost)/test?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open("mysql", "root:root@tcp(172.30.128.1:8080)/test") //DB connection
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/test") //DB connection
	if err != nil {
		log.Println("Could not connect to db")
		return err
	}
	db.DB().SetMaxIdleConns(2)
	db.DB().SetMaxOpenConns(2)

	db.DB().SetConnMaxLifetime(30 * time.Second)

	DB = db
	return nil
}
