package services

import (
	"../models"
	"../database"
	"../errors"
	"../utils"
	"github.com/imdario/mergo"
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

func (service RegistrationService) UpdateAttendee(attendee_update *models.Attendee, id int) {
	attendee := RegistrationServices.GetAttendeeByID(id)
	attendee.Collaborators = nil
	attendee.LongForms = nil
	attendee.ExtraInfos = nil

	if err := mergo.Merge(attendee, *attendee_update, mergo.WithOverride, mergo.WithTransformers(utils.TimeTransfomer{})); err != nil {
		panic(errors.UnprocessableError("Could not update attendee"))
	}

	tx := database.DB.Begin()

	if err := tx.Where(map[string]interface{}{"attendee_id": attendee.ID}).Delete(models.AttendeeLongForm{}).Error; err != nil {
		tx.Rollback()
		panic(errors.UnprocessableError("Could not update attendee"))
	}
	if err := tx.Where(map[string]interface{}{"attendee_id": attendee.ID}).Delete(models.AttendeeExtraInfo{}).Error; err != nil {
		tx.Rollback()
		panic(errors.UnprocessableError("Could not update attendee"))
	}
	if err := tx.Where(map[string]interface{}{"attendee_id": attendee.ID}).Delete(models.AttendeeCollaborator{}).Error; err != nil {
		tx.Rollback()
		panic(errors.UnprocessableError("Could not update attendee"))
	}

	if err := tx.Save(&attendee).Error; err != nil {
		tx.Rollback()
		panic(errors.UnprocessableError("Could not update attendee save error"))
	}

	if err := tx.Commit().Error; err != nil {
		panic(errors.UnprocessableError("Could not update attendee"))
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
