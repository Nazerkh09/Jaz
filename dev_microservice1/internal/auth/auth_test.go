package auth

import (
	"context"
	"testing"

	pb "github.com/Nazerkh09/jaz/dev_microservice1/api/auth"
	"github.com/stretchr/testify/assert"
)

func TestAuthService_RegisterUser(t *testing.T) {
	service := &AuthService{}

	req := &pb.RegisterUserRequest{
		Username: "testuser",
		Password: "testpassword",
	}

	res, err := service.RegisterUser(context.Background(), req)

	assert.Nil(t, err)
	assert.True(t, res.Success)
	assert.Equal(t, "User registered successfully", res.Message)
}

// Implement tests for other methods
