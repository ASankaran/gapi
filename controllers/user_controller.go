package controllers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/justinas/alice"
    "../services"
    "../models"
    "../middleware"
    "strconv"
)

type UserController struct {
    router *mux.Router
}

func (controller UserController) SetupRouter(route *mux.Route) {
    userRouter := route.Subrouter()

    userRouter.Handle("/", alice.New().ThenFunc(getCurrentUser)).Methods("GET")
    userRouter.Handle("/{id}", alice.New().ThenFunc(getUser)).Methods("GET")
    userRouter.Handle("/", alice.New(middleware.AuthMiddleware).ThenFunc(createUser)).Methods("POST")
}

func getCurrentUser(w http.ResponseWriter, r *http.Request) {
	user := services.UserServices.GetUserByID(1)
	json.NewEncoder(w).Encode(user)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	user := services.UserServices.GetUserByID(id)
	json.NewEncoder(w).Encode(user)

}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	services.UserServices.CreateUser(&user)
	json.NewEncoder(w).Encode("User was created")
}
