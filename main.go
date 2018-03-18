package main

import (
    "log"
    "net/http"
    "./controllers"
    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()
    controllers.SetupRouter(router)
    log.Fatal(http.ListenAndServe(":8000", router))
}
