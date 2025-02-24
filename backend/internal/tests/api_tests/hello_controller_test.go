package api_tests

import (
	"net/http"
	"testing"

	api_clients "github.com/KeishiIrisa/backend-go-template/internal/tests/clients"
)

func TestGetHelloWorld(t *testing.T) {
	// Set up the Gin router using the reusable function
	client := api_clients.NewTestClient(false)

	// Perform a request to the route
	response := client.PerformRequest("GET", "/", nil, nil)
	// Assert the response
	api_clients.AssertResponse(t, response, http.StatusOK, `"This server is running!"`)
}
