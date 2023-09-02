package services

import (
	"quickcooks/user-management/models"
	"quickcooks/user-management/repositories"
)

// A AuthenticationService is a provider for user authentication functionality
type AuthenticationService struct {
	userRepository repositories.IUserRepository
}

// NewAuthenticationService creates a new AuthenticationService instance with the
// given user repository
func NewAuthenticationService(
	userRepository repositories.IUserRepository,
) (*AuthenticationService, error) {
	authenticationService := &AuthenticationService{
		userRepository: userRepository,
	}

	return authenticationService, nil
}

// RegisterUser creates a new QuickCooks user with the given information
func (s *AuthenticationService) RegisterUser(name string, email string, password string) (*models.User, error) {
	user := &models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	return s.userRepository.Create(user)
}

// RegisterUser removes a new QuickCooks user with the given ID
func (s *AuthenticationService) UnregisterUser(userID uint) (*models.User, error) {
	user, err := s.userRepository.GetByID(userID)
	if err != nil {
		return nil, err
	}
	return s.userRepository.Delete(user)
}

// CheckUserEmailExists validates whether an a user with the given email exists
// in the database
func (s *AuthenticationService) CheckUserEmailExists(email string) bool {
	return s.userRepository.Exists(email)
}
