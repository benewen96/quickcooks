package services

import (
	"quickcooks/user-management/models"
	"quickcooks/user-management/repositories"
)

// A MyProfileService is a provider for user profile functionality
type MyProfileService struct {
	userRepository repositories.IUserRepository
}

// NewMyProfileService creates a new MyProfileService instance with the given
// user repository
func NewMyProfileService(userRepository repositories.IUserRepository) *MyProfileService {
	return &MyProfileService{
		userRepository: userRepository,
	}
}

// GetUserById returns the user with the given ID, if it exists
func (s *MyProfileService) GetUserByID(ID uint) (*models.User, error) {
	return s.userRepository.GetByID(ID)
}

func (s *MyProfileService) GetUserByEmail(email string) (*models.User, error) {
	return s.userRepository.GetByEmail(email)
}

// UpdateUserName updates the name of the user with the given ID, if it exists,
// to the given name
func (s *MyProfileService) UpdateUserName(userID uint, name string) (*models.User, error) {
	user, err := s.userRepository.GetByID(userID)
	if err != nil {
		return nil, err
	}
	return s.userRepository.UpdateName(user, name)
}

// UpdateUserEmail updates the email of the user with the given ID, if it
// exists, to the given email
func (s *MyProfileService) UpdateUserEmail(userID uint, email string) (*models.User, error) {
	user, err := s.userRepository.GetByID(userID)
	if err != nil {
		return nil, err
	}
	return s.userRepository.UpdateEmail(user, email)
}

// TODO: Password hashing

// UpdateUserPassword updates the password of the user with the given ID, if it
// exists, to the given password
func (s *MyProfileService) UpdateUserPassword(userID uint, password string) (*models.User, error) {
	user, err := s.userRepository.GetByID(userID)
	if err != nil {
		return nil, err
	}
	return s.userRepository.UpdatePassword(user, password)
}
