package models

import (
    "github.com/jinzhu/gorm"
)

type Stat struct {
	gorm.Model
	Category         string            `json:"Category,omitempty"                  sql:"not null;type:ENUM('REGISTRATION')"`
	Attribute        string            `json:"attribute,omitempty"                 sql:"not null"`
	Field            string            `json:"field,omitempty"                     sql:"not null"`
	Count            int               `json:"count,omitempty"                     sql:"not null;default:0"`
}
