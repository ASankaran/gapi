package middleware

import (
    "net/http"
    "fmt"
)

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    	fmt.Println("Hit auth middleware")
        next.ServeHTTP(w, r)
    })
}
