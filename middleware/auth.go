package middleware

import (
    "net/http"
    "../services"
    "../errors"
    "github.com/justinas/alice"
)

func AuthMiddleware(required string) alice.Constructor {
	return func(next http.Handler) http.Handler {
	    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	    	token := r.Header.Get("Authorization")
	    	if services.AuthServices.VerifyToken(token) {
	    		id := services.AuthServices.GetUserID(token)
	    		if services.AuthServices.HasRequiredAuth(id, required) {
	    			next.ServeHTTP(w, r)
	    			return
	    		}
	    	}
	    	panic(errors.UnauthorizedError("Bad auth token"))
	    })
	}
}


