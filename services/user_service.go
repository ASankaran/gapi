package services

import (
	"../models"
	"../database"
)

type UserService struct {
}

func (service UserService) CreateUser(user *models.User) {
	database.DB.Create(user)
}

func (service UserService) GetUserByID(id int) *models.User {
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return nil
	}
	return &user
}
