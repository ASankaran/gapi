package services

import (
	"../models"
	"../database"
	"../errors"
	"../utils"
)

type UserService struct {
}

func (service UserService) CreateUser(user *models.User, role string) {
	password, err := utils.Hash(user.Password)
	if err != nil {
		panic(errors.UnprocessableError("Could not create user"))
	}
	user.Password = password

	tx := database.DB.Begin()
	if tx.Error != nil {
		panic(errors.UnprocessableError("Could not create user"))
	}

	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		panic(errors.UnprocessableError("Could not create user"))
	}

	var user_role models.UserRole
	user_role.Role = role
	user_role.User = *user
	if err := tx.Create(&user_role).Error; err != nil {
		tx.Rollback()
		panic(errors.UnprocessableError("Could not create user"))
	}

	if err := tx.Commit().Error; err != nil {
		panic(errors.UnprocessableError("Could not create user"))
	}
}

func (service UserService) GetUserByID(id int) *models.User {
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		panic(errors.UnprocessableError("No matching id found"))
	}
	return &user
}

func (service UserService) GetUserByEmailPassword(email string, password string) *models.User {
	var user models.User
	if err := database.DB.Where(map[string]interface{}{"Email": email}).First(&user).Error; err != nil {
		panic(errors.UnprocessableError("No matching email and password found"))
	}
	if !utils.VerifyHash(password, user.Password) {
		panic(errors.UnprocessableError("No matching email and password found"))
	}
	return &user
}
