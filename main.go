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
	for _, r := range routes {
		http.Handle(r.path, middleware.Process(r.handler, append(defaultMiddlewares, r.extraMiddlewares...)))
	}
	log.Fatal(http.ListenAndServe(":8080", nil))
}
