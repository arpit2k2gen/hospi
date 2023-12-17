package main

import (
	"hospi/datastore"
	"hospi/handler"

	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	s := datastore.New()
	h := handler.New(s)

	app.GET("/patient/{id}", h.GetByID)
	app.POST("/patient", h.Create)
	app.PUT("/patient/{id}", h.Update)
	app.DELETE("/patient/{id}", h.Delete)

	// starting the server on a custom port
	app.Server.HTTP.Port = 8000
	app.Start()
}
