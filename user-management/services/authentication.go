package services

import (
	"quickcooks/user-management/models"
	"quickcooks/user-management/repositories"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticationServiceConfig struct {
	JwtSecret string
}

// A AuthenticationService is a provider for user authentication functionality
type AuthenticationService struct {
	userRepository repositories.IUserRepository
	config         AuthenticationServiceConfig
}

// NewAuthenticationService creates a new AuthenticationService instance with the
// given user repository
func NewAuthenticationService(
	userRepository repositories.IUserRepository,
	config AuthenticationServiceConfig,
) (*AuthenticationService, error) {
	authenticationService := &AuthenticationService{
		userRepository: userRepository,
		config:         config,
	}

	return authenticationService, nil
}

// RegisterUser creates a new QuickCooks user with the given information
func (s *AuthenticationService) RegisterUser(
	name string, email string, password string,
) (
	*models.User, error,
) {
	hashedPassword, err := s.hashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}
	return s.userRepository.Create(user)
}

func (s *AuthenticationService) Login(
	email string, password string,
) (
	token string, authed bool, err error,
) {
	user, err := s.userRepository.GetByEmail(email)
	if err != nil {
		return "", false, err
	}

	hashedPassword, err := s.hashPassword(password)
	if err != nil {
		return "", false, err
	}

	if !checkHashedPassword(user.Password, hashedPassword) {
		return "", false, nil
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":  user.Name,
		"email": user.Email,
		"roles": user.RoleAssignments,
	})

	tokenString, err := jwtToken.SignedString(s.config.JwtSecret)
	if err != nil {
		return "", false, err
	}

	return tokenString, false, nil
}

// UnregisterUser removes a new QuickCooks user with the given ID
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

func (s *AuthenticationService) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkHashedPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
