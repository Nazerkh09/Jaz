package user

import (
	"encoding/json"
	"net/http"
)

// UserHandler represents the HTTP handler for user-related endpoints.
type UserHandler struct {
	userService *UserService
}

// NewUserHandler creates a new UserHandler with the given UserService.
func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{userService: service}
}

// GetUserByIDHandler handles the GET /api/user/{id} endpoint.
func (h *UserHandler) GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")

	user, err := h.userService.GetUserByID(userID)
	if err != nil {
		http.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
