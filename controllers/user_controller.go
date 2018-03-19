package controllers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/justinas/alice"
    "../services"
    "../models"
    "../middleware"
    "../requests"
    "../responses"
    "strconv"
)

type UserController struct {
    router *mux.Router
}

func (controller UserController) SetupRouter(route *mux.Route) {
    userRouter := route.Subrouter()

    userRouter.Handle("/", alice.New(middleware.AuthMiddleware).ThenFunc(getCurrentUser)).Methods("GET")
    userRouter.Handle("/{id}", alice.New(middleware.AuthMiddleware).ThenFunc(getUser)).Methods("GET")
    userRouter.Handle("/", alice.New().ThenFunc(createUser)).Methods("POST")
    userRouter.Handle("/login", alice.New().ThenFunc(loginUser)).Methods("POST")
}

func getCurrentUser(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	id := services.AuthServices.GetUserID(token)
	user := services.UserServices.GetUserByID(id)
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
	json.NewEncoder(w).Encode(user)
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	var request requests.Login
	json.NewDecoder(r.Body).Decode(&request)
	user := services.UserServices.GetUserByEmailPassword(request.Email, request.Password)
	if user == nil {
		json.NewEncoder(w).Encode(nil)
		return
	}
	token := services.AuthServices.GenerateToken(user)
	json.NewEncoder(w).Encode(responses.Token{Token: token})
}
