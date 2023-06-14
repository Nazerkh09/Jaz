package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Nazerkh09/jaz/dev_microservice1/internal/auth"
)

func TestHandleRegister(t *testing.T) {
	// Create a registration request
	registrationRequest := auth.RegistrationRequest{
		Username: "testuser",
		Password: "testpassword",
		Email:    "test@example.com",
	}

	// Encode the request body
	requestBody, err := json.Marshal(registrationRequest)
	if err != nil {
		t.Fatalf("failed to encode request body: %v", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "/api/register", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Call the handleRegister handler function
	handleRegister(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rr.Code)
	}

	// Perform additional assertions or checks based on your requirements
}

// Implement other test functions for handleLogin, handleValidateToken, and handleGetUserPermissions
