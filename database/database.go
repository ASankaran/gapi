package database

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
    "../models"
)

var DB, _ = gorm.Open("sqlite3", "test.db")

func Migrate() {
	DB.AutoMigrate(&models.User{})
}
