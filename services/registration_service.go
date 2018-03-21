package services

import (
	"../models"
	"../database"
	"../errors"
)

type RegistrationService struct {
}

func (service RegistrationService) CreateAttendee(attendee *models.Attendee) {
	tx := database.DB.Begin()
	if tx.Error != nil {
		panic(errors.UnprocessableError("Could not register attendee"))
	}

	if err := tx.Create(attendee).Error; err != nil {
		tx.Rollback()
		panic(errors.UnprocessableError("Could not register attendee"))
	}

	user := UserServices.GetUserByID(attendee.UserID)
	var user_role models.UserRole
	user_role.Role = "ATTENDEE"
	user_role.User = *user
	if err := tx.Create(&user_role).Error; err != nil {
		tx.Rollback()
		panic(errors.UnprocessableError("Could not register attendee"))
	}

	if err := tx.Commit().Error; err != nil {
		panic(errors.UnprocessableError("Could not register attendee"))
	}
}

func (service RegistrationService) GetAttendeeByID(id int) *models.Attendee {
	var attendee models.Attendee
	if err := database.DB.Where(map[string]interface{}{"user_id": id}).First(&attendee).Error; err != nil {
		panic(errors.UnprocessableError("No matching id found"))
	}
	database.DB.Preload("Collaborators").Find(&attendee)
	database.DB.Preload("LongForms").Find(&attendee)
	database.DB.Preload("ExtraInfos").Find(&attendee)
	return &attendee
}
