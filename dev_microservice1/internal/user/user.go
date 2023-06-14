package user

import "github.com/Nazerkh09/jaz/dev_microservice1/api/user"

// UserRepository provides an interface to interact with the user data store.
type UserRepository interface {
	GetUserByID(userID string) (*user.User, error)
}

// UserService represents the user service.
type UserService struct {
	userRepo UserRepository
}

// NewUserService creates a new UserService with the given UserRepository.
func NewUserService(repo UserRepository) *UserService {
	return &UserService{userRepo: repo}
}

// GetUserByID retrieves a user by their ID.
func (s *UserService) GetUserByID(userID string) (*user.User, error) {
	return s.userRepo.GetUserByID(userID)
}
