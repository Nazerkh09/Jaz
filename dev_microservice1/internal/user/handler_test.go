package user

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Nazerkh09/jaz/dev_microservice1/api/user"
	"github.com/stretchr/testify/assert"
)

type MockUserService struct{}

func (s *MockUserService) GetUserByID(userID string) (*user.User, error) {
	// Mock implementation
	return &user.User{ID: userID, Name: "John Doe"}, nil
}

func TestGetUserByIDHandler(t *testing.T) {
	service := &MockUserService{}
	handler := NewUserHandler(service)

	req, err := http.NewRequest("GET", "/api/user?id=123", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	recorder := httptest.NewRecorder()

	handler.GetUserByIDHandler(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	expectedUser := &user.User{ID: "123", Name: "John Doe"}
	var responseBody user.User

	err = json.Unmarshal(recorder.Body.Bytes(), &responseBody)
	if err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	assert.Equal(t, expectedUser, &responseBody)
}
