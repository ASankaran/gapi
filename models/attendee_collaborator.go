package models

import (
    "github.com/jinzhu/gorm"
)

type AttendeeCollaborator struct {
	gorm.Model
	Collaborator          string     `json:"collaborator,omitempty"             sql:"not null"`
	AttendeeID            int
}
