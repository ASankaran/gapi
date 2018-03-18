package main

import (
    "log"
    "net/http"
    "./controllers"
    "./database"
    "github.com/gorilla/mux"
)

func main() {
	database.Migrate()
    router := mux.NewRouter()
    controllers.SetupRouter(router)
    log.Fatal(http.ListenAndServe(":8000", router))
}
