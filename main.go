package main

import (
    "./controllers"
    "./database"
    "./middleware"
    "github.com/gorilla/mux"

    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
)

var initialized = false
var gorillaMuxAdapter *gorillamux.GorillaMuxAdapter

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

    if !initialized {
        database.Migrate()
        router := mux.NewRouter()
        controllers.SetupRouter(router)
        
        router.Use(middleware.ErrorMiddleware)
        router.Use(middleware.ContentTypeMiddleware)

        gorillaMuxAdapter = gorillamux.New(router)
        initialized = true
    }

    return gorillaMuxAdapter.Proxy(req)
}

func main() {
    lambda.Start(Handler)
}
