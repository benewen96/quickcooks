package services

import (
	"quickcooks/user-management/models"
	"quickcooks/user-management/repositories"
)

// A RegistrationService is a provider for user registration functionality
type RegistrationService struct {
	userRepository repositories.IUserRepository
}

// NewRegistrationService creates a new RegistrationService instance with the
// given user repository
func NewRegistrationService(userRepository repositories.IUserRepository) *RegistrationService {
	return &RegistrationService{
		userRepository: userRepository,
	}
}

// TODO: Email verification
// TODO: Password hashing

// RegisterUser creates a new QuickCooks user with the given information
func (s *RegistrationService) RegisterUser(name string, email string, password string) (*models.User, error) {
	user := &models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	return s.userRepository.Create(user)
}

// RegisterUser removes a new QuickCooks user with the given ID
func (s *RegistrationService) UnregisterUser(userID uint) (*models.User, error) {
	user, err := s.userRepository.GetByID(userID)
	if err != nil {
		return nil, err
	}
	return s.userRepository.Delete(user)
}
