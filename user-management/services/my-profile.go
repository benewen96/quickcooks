package services

import (
	"quickcooks/user-management/models"
	"quickcooks/user-management/repositories"
)

type MyProfileService struct {
	userRepository repositories.IUserRepository
}

func NewMyProfileService(userRepository repositories.IUserRepository) *MyProfileService {
	return &MyProfileService{
		userRepository: userRepository,
	}
}

func (s *MyProfileService) GetUserByID(ID uint) (*models.User, error) {
	return s.userRepository.GetByID(ID)
}

func (s *MyProfileService) UpdateUserName(userID uint, name string) (*models.User, error) {
	user, err := s.userRepository.GetByID(userID)
	if err != nil {
		return nil, err
	}
	return s.userRepository.UpdateName(user, name)
}

func (s *MyProfileService) UpdateUserEmail(userID uint, email string) (*models.User, error) {
	user, err := s.userRepository.GetByID(userID)
	if err != nil {
		return nil, err
	}
	return s.userRepository.UpdateEmail(user, email)
}

// TODO: Password hashing

func (s *MyProfileService) UpdateUserPassword(userID uint, password string) (*models.User, error) {
	user, err := s.userRepository.GetByID(userID)
	if err != nil {
		return nil, err
	}
	return s.userRepository.UpdatePassword(user, password)
}
