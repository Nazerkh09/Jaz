package user

import (
	"testing"

	"github.com/Nazerkh09/jaz/dev_microservice1/api/user"
	"github.com/stretchr/testify/assert"
)

type MockUserRepository struct{}

func (r *MockUserRepository) GetUserByID(userID string) (*user.User, error) {
	// Mock implementation
	return &user.User{Id: userID, Name: "John Doe"}, nil
}

func TestGetUserByID(t *testing.T) {
	repo := &MockUserRepository{}
	service := NewUserService(repo)

	// Test case
	userID := "123"
	expectedUser := &user.User{Id: userID, Name: "John Doe"}

	result, err := service.GetUserByID(userID)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, result)
}
