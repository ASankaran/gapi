package models

import (
	"time"
    "github.com/jinzhu/gorm"
)

type Attendee struct {
	gorm.Model
	User                      User       `gorm:"foreignkey:UserID"`
	UserID                    int
	ShirtSize                 string     `json:"shirtsize,omitempty"               sql:"not null;type:ENUM('S', 'M', 'L', 'XL')"`
	Diet                      string     `json:"diet,omitempty"                    sql:"not null;type:ENUM('NONE', 'VEGETARIAN', 'VEGAN', 'GLUTEN_FREE');default:'NONE'"`
	Age                       int        `json:"age,omitempty"                     sql:"not null"`
	GraduateYear              int        `json:"graduationyear,omitempty"          sql:"not null"`
	Transportation            string     `json:"transportation,omitempty"          sql:"not null;type:ENUM('NOT_NEEDED', 'BUS_REQUESTED', 'IN_STATE', 'OUT_OF_STATE', 'INTERNATIONAL');default:'NOT_NEEDED'"`
	School                    string     `json:"school,omitempty"                  sql:"not null"`
	Major                     string     `json:"major,omitempty"                   sql:"not null"`
	Gender                    string     `json:"gender,omitempty"                  sql:"not null;type:ENUM('MALE', 'FEMALE', 'NON_BINARY', 'OTHER')"`
	ProfessionalInterest      string     `json:"professionalinterest,omitempty"    sql:"not null;type:ENUM('NONE', 'INTERNSHIP', 'FULLTIME', 'BOTH');default:'NONE'"`
	Github                    string     `json:"github,omitempty"                  sql:"not null"`
	Linkedin                  string     `json:"linkedin,omitempty"                sql:"not null"`
	Interests                 string     `json:"interests,omitempty"               sql:"null;default:null"`
	Status                    string     `json:"status,omitempty"                  sql:"not null;type:ENUM('ACCEPTED', 'WAITLISTED', 'REJECTED', 'PENDING');default:'PENDING'"`
	IsNovice                  bool       `json:"isnovice,omitempty"                sql:"not null;default:0"`
	IsPrivate                 bool       `json:"isprivate,omitempty"               sql:"not null;default:0"`
	HasLightningInterest      bool       `json:"haslightninginterest,omitempty"    sql:"not null;default:0"`
	PhoneNumber               string     `json:"phonenumber,omitempty"             sql:"null;default:null"`
	Priority                  bool       `json:"priority,omitempty"                sql:"not null;default:0"`
	Wave                      int        `json:"wave,omitempty"                    sql:"not null;default:0"`
	Reviewer                  string     `json:"reviewer,omitempty"                sql:"null;default:null"`
	ReviewTime                time.Time  `json:"reviewertime,omitempty"            sql:"null;default:null"`
}
