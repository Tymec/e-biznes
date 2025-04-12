package lib

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var database *gorm.DB
var err error

func InitDB() {
	database, err = gorm.Open(sqlite.Open("db.sqlite3"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return database
}
