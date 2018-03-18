package main

import (
    "log"
    "net/http"
    "./controllers"
    "./database"
    "./middleware"
    "github.com/gorilla/mux"
)

func main() {
	database.Migrate()
    router := mux.NewRouter()
    controllers.SetupRouter(router)
    
    router.Use(middleware.ContentTypeMiddleware)
    log.Fatal(http.ListenAndServe(":8000", router))
}
