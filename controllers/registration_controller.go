package controllers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/justinas/alice"
    "../middleware"
    "../services"
    "../requests"
    "strconv"
)

type RegistrationController struct {
    router *mux.Router
}

func (controller RegistrationController) SetupRouter(route *mux.Route) {
    registrationRouter := route.Subrouter()

    registrationRouter.Handle("/", alice.New(middleware.AuthMiddleware(services.AUTH_ATTENDEE)).ThenFunc(getCurrentAttendee)).Methods("GET")
    registrationRouter.Handle("/{id}", alice.New(middleware.AuthMiddleware(services.AUTH_ADMIN)).ThenFunc(getAttendee)).Methods("GET")
    registrationRouter.Handle("/", alice.New(middleware.AuthMiddleware(services.AUTH_USER)).ThenFunc(createAttendee)).Methods("POST")
}

func getCurrentAttendee(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	id := services.AuthServices.GetUserID(token)
	attendee := services.RegistrationServices.GetAttendeeByID(id)
	json.NewEncoder(w).Encode(attendee)
}

func getAttendee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	attendee := services.RegistrationServices.GetAttendeeByID(id)
	json.NewEncoder(w).Encode(attendee)

}

func createAttendee(w http.ResponseWriter, r *http.Request) {
	var attendee_registration requests.AttendeeRegistration
	json.NewDecoder(r.Body).Decode(&attendee_registration)

	attendee := attendee_registration.Attendee

	token := r.Header.Get("Authorization")
	id := services.AuthServices.GetUserID(token)
	user := services.UserServices.GetUserByID(id)
	attendee.User = *user

	services.RegistrationServices.CreateAttendee(&attendee)
	json.NewEncoder(w).Encode(attendee)
}
