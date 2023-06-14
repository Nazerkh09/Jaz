package auth

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	pb "github.com/Nazerkh09/jaz/dev_microservice1/api/auth"
)

// MockAuthService is a mock implementation of the AuthServiceServer interface.
type MockAuthService struct{}

// RegisterUser is a mock implementation of the RegisterUser method.
func (s *MockAuthService) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	// Mock implementation
	return &pb.RegisterUserResponse{
		UserId: "123",
	}, nil
}

// TestHandleRegister tests the handleRegister function.
func TestHandleRegister(t *testing.T) {
	// Create a new instance of the handler
	handler := NewAuthHandler(&MockAuthService{})

	// Create a new request with the necessary data
	req, err := http.NewRequest("POST", "/api/register", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	// Create a response recorder to record the response
	recorder := httptest.NewRecorder()

	// Call the handleRegister function
	handler.handleRegister(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, recorder.Code)
	}

	// Check the response body
	expectedBody := "User registered successfully"
	if recorder.Body.String() != expectedBody {
		t.Errorf("expected body %q, got %q", expectedBody, recorder.Body.String())
	}
}
