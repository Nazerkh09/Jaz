package main

import (
	"encoding/json"
	"log"
	"net/http"

	auth "github.com/Nazerkh09/jaz/dev_microservice1/internal/auth"
)

func main() {
	go auth.RunGRPCServer()

	router := http.NewServeMux()

	router.HandleFunc("/api/register", handleRegister)
	router.HandleFunc("/api/login", handleLogin)
	router.HandleFunc("/api/validate-token", handleValidateToken)
	router.HandleFunc("/api/user-permissions", handleGetUserPermissions)

	log.Println("Starting HTTP/REST server on port 8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// Implement HTTP request handlers for each API endpoint

func handleRegister(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var registrationRequest auth.RegistrationRequest
	err := json.NewDecoder(r.Body).Decode(&registrationRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Perform registration logic
	err = auth.RegisterUser(registrationRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var loginRequest auth.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Perform login logic
	token, err := auth.Login(loginRequest)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Return token in the response
	response := struct {
		Token string `json:"token"`
	}{
		Token: token,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func handleValidateToken(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var tokenRequest auth.TokenRequest
	err := json.NewDecoder(r.Body).Decode(&tokenRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Perform token validation logic
	valid, err := auth.ValidateToken(tokenRequest.Token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return validation result in the response
	response := struct {
		Valid bool `json:"valid"`
	}{
		Valid: valid,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func handleGetUserPermissions(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var userRequest auth.UserRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get user permissions
	permissions, err := auth.GetUserPermissions(userRequest.UserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return user permissions in the response
	response := struct {
		Permissions []string `json:"permissions"`
	}{
		Permissions: permissions,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
