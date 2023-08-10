package services

import (
	"quickcooks/user-management/models"
	"quickcooks/user-management/repositories"
)

type RegistrationService struct {
	UserRepository repositories.IUserRepository
}

func NewRegistrationService(userRepository repositories.IUserRepository) *RegistrationService {
	return &RegistrationService{
		UserRepository: userRepository,
	}
}

// TODO: Email verification
// TODO: Password hashing

func (s *RegistrationService) RegisterUser(name string, email string, password string) (*models.User, error) {
	user := &models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	return s.UserRepository.Create(user)
}

func (s *RegistrationService) UnregisterUser(userID uint) (*models.User, error) {
	user, err := s.UserRepository.GetByID(userID)
	if err != nil {
		return nil, err
	}
	return s.UserRepository.Delete(user)
}
