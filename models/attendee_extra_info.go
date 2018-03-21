package models

import (
    "github.com/jinzhu/gorm"
)

type AttendeeExtraInfo struct {
	gorm.Model
	Website               string     `json:"website,omitempty"             sql:"not null"`
	AttendeeID            int
}
