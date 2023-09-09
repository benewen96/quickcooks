package authentication

import (
	"quickcooks/user-management/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AuthenticationTestSuite struct {
	suite.Suite
	AuthenticationService *AuthenticationService
}

func (suite *AuthenticationTestSuite) SetupTest() {
	suite.AuthenticationService, _ = NewAuthenticationService(repositories.NewMockUserRepository(), AuthenticationServiceConfig{
		JwtSecret: "mockSecret!",
	})
}

func (suite *AuthenticationTestSuite) TestPasswordHashMatch() {
	password := "password"
	hash, _ := suite.AuthenticationService.HashPassword(password)
	result := suite.AuthenticationService.CheckHashedPassword(password, hash)

	assert.Equal(suite.T(), true, result)
}

func (suite *AuthenticationTestSuite) TestPasswordHashMismatch() {
	password := "password"
	hash, _ := suite.AuthenticationService.HashPassword(password)
	result := suite.AuthenticationService.CheckHashedPassword("incorrect_password", hash)

	assert.Equal(suite.T(), false, result)
}

func (suite *AuthenticationTestSuite) TestSuccessfulLogin() {
	email := "test@test.com"
	plainPassword := "p4ssw0rd!"

	_, err := suite.AuthenticationService.RegisterUser("Mr Test", email, plainPassword)
	assert.Nil(suite.T(), err)

	_, _, authed, err := suite.AuthenticationService.Login(email, plainPassword)
	assert.Nil(suite.T(), err)
	assert.True(suite.T(), authed)
}

func TestAuthenticationTestSuite(t *testing.T) {
	suite.Run(t, new(AuthenticationTestSuite))
}
