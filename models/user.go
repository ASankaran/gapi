package models

import (
    "github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
    Firstname string   `json:"firstname,omitempty"    sql:"not null"`
    Lastname  string   `json:"lastname,omitempty"     sql:"not null"`
    Email     string   `json:"email,omitempty"        sql:"not null"`
    Password  string   `json:"password,omitempty"     sql:"not null"`
}
