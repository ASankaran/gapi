package services

import (
	"../models"
	"../database"
	"../errors"
	"../utils"
)

type UserService struct {
}

func (service UserService) CreateUser(user *models.User) {
	password, err := utils.Hash(user.Password)
	if err != nil {
		panic(errors.UnprocessableError("Could not create user"))
	}
	user.Password = password
	database.DB.Create(user)
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
