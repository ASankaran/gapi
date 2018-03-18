package middleware

import (
    "net/http"
)

// func ContentTypeMiddleware(next http.Handler) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//     	w.Header().Set("Content-Type", "application/json")
//         next.ServeHTTP(w, r)
//     })
// }

func ContentTypeMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Header().Set("Content-Type", "application/json")
	next(w, r)
}