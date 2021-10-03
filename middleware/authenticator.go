package middleware

import (
	"log"
	"net/http"
)

func Authenticator() Middleware {
	return middlewareFunc(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("Authenticator Middleware")
			if r.Header.Get("X-Token") == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	})
}
