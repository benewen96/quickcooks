package services

import "quickcooks/user-management/repositories"

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

// CheckUserEmailExists validates whether an a user with the given email exists
// in the database
func (s *AuthenticationService) CheckUserEmailExists(email string) bool {
	return s.userRepository.Exists(email)
}
