package services

import (
	"../models"
	"../database"
)

type UserService struct {
}

func (service UserService) CreateUser(user *models.User) {
	database.DB.AutoMigrate(&models.User{})
	database.DB.Create(user)
}

func (service UserService) GetUserByID(id int) models.User {
	var user models.User
	database.DB.First(&user, id)
	return user
}
