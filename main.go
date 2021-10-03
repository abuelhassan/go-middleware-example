package main

import (
	"log"
	"net/http"

	"github.com/abuelhassan/go-middleware-example/handler"
	"github.com/abuelhassan/go-middleware-example/middleware"
)

type route struct {
	path             string
	handler          http.HandlerFunc
	extraMiddlewares []middleware.Middleware
}

var (
	defaultMiddlewares = []middleware.Middleware{middleware.Logger()}

	routes = []route{
		{
			path:             "/employees",
			handler:          handler.GetEmployees,
			extraMiddlewares: nil,
		},
		{
			path:             "/clients",
			handler:          handler.GetClients,
			extraMiddlewares: []middleware.Middleware{middleware.Authenticator()},
		},
	}
)

func main() {
	mux := http.NewServeMux()
	for _, r := range routes {
		mux.Handle(r.path, middleware.Process(r.handler, r.extraMiddlewares...))
	}
	wrappedMux := middleware.Process(mux, defaultMiddlewares...)
	log.Fatal(http.ListenAndServe(":8080", wrappedMux))
}
