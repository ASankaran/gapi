package models

import (
    "github.com/jinzhu/gorm"
)

type UserRole struct {
	gorm.Model
	User        User     `gorm:"foreignkey:UserID"`
	UserID      int
	Role        string   `sql:"not null;type:ENUM('ADMIN', 'USER')"`
}
