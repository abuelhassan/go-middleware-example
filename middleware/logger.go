package middleware

import (
	"context"
	"log"
	"net/http"
	"os"
)

func Logger() Middleware {
	return middlewareFunc(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("Logger Middleware")
			l := log.New(os.Stdout, "rid:123 ", log.Ldate|log.Lmicroseconds|log.LUTC)
			r = r.WithContext(context.WithValue(r.Context(), "logger", l))
			next.ServeHTTP(w, r)
		})
	})
}
