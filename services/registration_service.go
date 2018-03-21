package services

import (
	"../models"
	"../database"
	"../errors"
)

type RegistrationService struct {
}

func (service RegistrationService) CreateAttendee(attendee *models.Attendee) {
	if err := database.DB.Create(attendee).Error; err != nil {
		panic(errors.UnprocessableError("Could not register attendee"))
	}

	user := UserServices.GetUserByID(attendee.UserID)
	UserServices.AddUserRole(user, "ATTENDEE")
}

func (service RegistrationService) GetAttendeeByID(id int) *models.Attendee {
	var attendee models.Attendee
	if err := database.DB.Where(map[string]interface{}{"user_id": id}).First(&attendee).Error; err != nil {
		panic(errors.UnprocessableError("No matching id found"))
	}
	return &attendee
}
