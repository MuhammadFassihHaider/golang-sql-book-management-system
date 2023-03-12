package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func ConnectToDB() {
	dsn := "root:fassih123@tcp(127.0.0.1:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"

	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
