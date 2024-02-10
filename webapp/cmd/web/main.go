package main

import (
	"log"
	"net/http"
)

type application struct{}

func main() {
	// set up an app config
	app := application{}

	// get application routes
	mux := app.routes()

	// register routes
	log.Println("Starting server on port 8080...")

	// static asseets
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
