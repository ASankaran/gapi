package main

import (
    "log"
    "net/http"
    "./controllers"
    "./database"
    "./middleware"
    "github.com/gorilla/mux"
    "github.com/urfave/negroni"
)

func main() {
	database.Migrate()
    router := mux.NewRouter()
    controllers.SetupRouter(router)

    routerMiddlewre := negroni.Classic()
    routerMiddlewre.Use(negroni.HandlerFunc(middleware.ContentTypeMiddleware))
    routerMiddlewre.UseHandler(router)
    
    log.Fatal(http.ListenAndServe(":8000", routerMiddlewre))
}
