package models

import (
    "github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
    Firstname string   `json:"firstname,omitempty"`
    Lastname  string   `json:"lastname,omitempty"`
    Email     string   `json:"email,omitempty"`
}
