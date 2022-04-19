package main

import (
	"github.com/aerogo/aero"
)

func main() {
	app := aero.New()

	app.Use(RequestLog)

	// Add routes to the app
	router := app.Router()
	for _, route := range GetRoutings() {
		router.Add(route.method, route.path, route.handler)
	}

	app.Run()
}
