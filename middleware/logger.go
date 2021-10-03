package middleware

import (
	"log"
	"net/http"
)

func Logger() Middleware {
	return middlewareFunc(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("Logger Middleware")
			next.ServeHTTP(w, r)
		})
	})
}
