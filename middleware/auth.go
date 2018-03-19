package middleware

import (
    "net/http"
    "../services"
)

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    	token := r.Header.Get("Authorization")
    	if services.AuthServices.VerifyToken(token) {
    		next.ServeHTTP(w, r)
    	}
    })
}
