package api

import "net/http"

func RegisterEndpoints() {

	// Register the user endpoint
	http.HandleFunc("/user", GetUser)
}
