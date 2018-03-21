package models

import (
    "github.com/jinzhu/gorm"
)

type AttendeeLongForm struct {
	gorm.Model
	Info                  string     `json:"info,omitempty"             sql:"not null"`
	AttendeeID            int
}
