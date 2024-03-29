package controllers

import (
    "github.com/gorilla/mux"
)

type controller interface {
	SetupRouter(route *mux.Route)
}

func SetupRouter(router *mux.Router) {
	userController := UserController{}
	userController.SetupRouter(router.PathPrefix("/user"))
	registrationController := RegistrationController{}
	registrationController.SetupRouter(router.PathPrefix("/registration"))
	statsController := StatsController{}
	statsController.SetupRouter(router.PathPrefix("/stats"))
}
