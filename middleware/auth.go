package middleware

import (
    "net/http"
    "../services"
    "../errors"
)

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    	token := r.Header.Get("Authorization")
    	if services.AuthServices.VerifyToken(token) {
    		next.ServeHTTP(w, r)
    		return
    	}
    	panic(errors.UnauthorizedError("Bad auth token"))
    })
}
