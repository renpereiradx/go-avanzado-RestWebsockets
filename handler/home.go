package handler

import (
	"encoding/json"
	"net/http"

	"github.com/renpereiradx/go-avanzado-RestWebsocket/server"
)

// HomeResponse represents the response structure for the home endpoint.
type HomeResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

// HomeHandler handles the home route and returns a http.HandlerFunc.
// It takes a server.Server as a parameter and returns a function that handles the request.
// The function sets the Content-Type header to "application/json",
// writes a 200 status code to the response writer, and encodes a HomeResponse struct as JSON.
func HomeHandler(server server.Server) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.WriteHeader(http.StatusOK)
		json.NewEncoder(responseWriter).Encode(HomeResponse{
			Message: "Welcome to the API",
			Status:  true,
		})
	}
}
