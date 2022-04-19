package main

import (
	"github.com/aerogo/aero"
)

type route struct {
	handler aero.Handler
	method string
	path string
}

// Create a new route
func newRoute(method string, path string, handler aero.Handler) *route {
	result := route{ handler: handler, method: method, path: path }
	
	return &result
}

// Return all application routings
func GetRoutings() []*route {
	result := []*route {
		
		// Home
		newRoute("GET", "/", func(ctx aero.Context) error {
			return ctx.File("public/html/index.html")
		}),

		// Javascript
		newRoute("GET", "/js/*file", func(ctx aero.Context) error {
			return ctx.File("public/js/" + ctx.Get("file"))
		}),
	}
	
	return result
}
