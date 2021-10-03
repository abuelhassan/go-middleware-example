package middleware

import "net/http"

type Middleware interface {
	Apply(http.Handler) http.Handler
}

// Process calls all middlewares in the given order.
func Process(h http.Handler, ms ...Middleware) http.Handler {
	for i := len(ms) - 1; i >= 0; i-- {
		h = ms[i].Apply(h)
	}
	return h
}

type middlewareFunc func(http.Handler) http.Handler

func (f middlewareFunc) Apply(h http.Handler) http.Handler {
	return f(h)
}
