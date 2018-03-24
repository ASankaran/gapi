package database

import (
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
    "../models"
    "fmt"
)

var DB *gorm.DB

func init() {
	var db, err = gorm.Open("mysql", "root:pass123@tcp(172.17.0.1:3306)/gapi?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
	}

	DB = db;
}

func Migrate() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.UserRole{})
	DB.AutoMigrate(&models.Attendee{})
	DB.AutoMigrate(&models.AttendeeCollaborator{})
	DB.AutoMigrate(&models.AttendeeLongForm{})
	DB.AutoMigrate(&models.AttendeeExtraInfo{})
	DB.AutoMigrate(&models.Stat{})
}
