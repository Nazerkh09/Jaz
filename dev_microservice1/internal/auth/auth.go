package auth

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	pb "github.com/Nazerkh09/jaz/dev_microservice1/api/auth"
)

var mongoClient *mongo.Client

type AuthService struct{}

type RegistrationRequest struct {
	// Define the fields of the registration request
	// For example:
	Username string `json:"username"`
	Password string `json:"password"`
	// Add more fields as needed
}

// LoginRequest represents the login request data
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenRequest struct {
	Token string `json:"token"`
}

func ValidateToken(token string) (bool, error) {
	// Your validation logic here
	// Return true if the token is valid, otherwise false
	// Return an error if there is an issue with token validation

	// Example implementation:
	if token == "valid_token" {
		return true, nil
	}

	return false, nil
}

type UserRequest struct {
	UserID string `json:"userID"`
}

func GetUserPermissions(userID string) ([]string, error) {
	// TODO: Implement the logic to fetch user permissions from the database or any other source.
	// For now, let's return a sample list of permissions.
	permissions := []string{"read", "write", "delete"}
	return permissions, nil
}

func Login(request LoginRequest) (string, error) {
	// Create a new MongoDB collection instance
	collection := mongoClient.Database("microservice1").Collection("user")

	// Query the user document based on the provided username and password
	filter := bson.M{
		"username": request.Username,
		"password": request.Password,
	}
	var user bson.M
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate and return the access token
	token := "example-token"
	return token, nil
}

func RegisterUser(request RegistrationRequest) error {
	// Create a new MongoDB collection instance
	collection := mongoClient.Database("microservice1").Collection("user")

	// Create a user document
	user := bson.M{
		"username": request.Username,
		"password": request.Password,
	}

	// Insert the user document into the collection
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	// Implementation logic for registering a user
	return &pb.RegisterUserResponse{
		Success: true,
		Message: "User registered successfully",
	}, nil
}

func (s *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// Implementation logic for user login
	return &pb.LoginResponse{
		Success:     true,
		Message:     "Login successful",
		AccessToken: "sample-access-token", // Corrected field name
	}, nil
}

func (s *AuthService) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	// Implementation logic for validating the access token
	return &pb.ValidateTokenResponse{
		Success: true,
		Message: "Token validated successfully",
		Valid:   true,
	}, nil
}

func (s *AuthService) GetUserPermissions(ctx context.Context, req *pb.GetUserPermissionsRequest) (*pb.GetUserPermissionsResponse, error) {
	// Implementation logic for retrieving user permissions
	return &pb.GetUserPermissionsResponse{
		Success:     true,
		Message:     "Permissions retrieved successfully",
		Permissions: []string{"permission1", "permission2"},
	}, nil
}
