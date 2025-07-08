// Package declaration
package main

import (
	"fit/api"
	"fit/application"
	"fmt"
	"net/http"
)

// Entry point
func main() {

	// Start the application
	application.Start()

	// Register endpoints
	api.RegisterEndpoints()

	// Register the static file server
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Host
	err := http.ListenAndServe(":8082", nil)
	// err := http.ListenAndServeTLS(":8082", "ca/g10.app.crt", "ca/g10.app.key", nil)

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
