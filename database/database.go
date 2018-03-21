package database

import (
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
    "../models"
)

var DB, _ = gorm.Open("mysql", "root:pass123@/gapi?charset=utf8&parseTime=True&loc=Local")

func Migrate() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.UserRole{})
	DB.AutoMigrate(&models.Attendee{})
	DB.AutoMigrate(&models.AttendeeCollaborator{})
}
