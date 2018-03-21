package controllers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/justinas/alice"
    "../middleware"
    "../services"
    "../responses"
)

type StatsController struct {
    router *mux.Router
}

func (controller StatsController) SetupRouter(route *mux.Route) {
    statsRouter := route.Subrouter()

    statsRouter.Handle("/", alice.New(middleware.AuthMiddleware(services.AUTH_ADMIN)).ThenFunc(getStats)).Methods("GET")
}

func getStats(w http.ResponseWriter, r *http.Request) {
	var stats responses.AllStats

	stats.Registration = services.StatsServices.GetRegistrationStats()

	json.NewEncoder(w).Encode(stats)
}
