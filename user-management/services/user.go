package services

import (
	"quickcooks/user-management/db"
	"quickcooks/user-management/models"
)

type IUserService interface {
	CreateUser(name string, email string, password string) (models.User, error)
	GetUserByID(ID uint) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	GetUsersByTenantID(ID uint) ([]models.User, error)
}

type UserService struct{}

func (us *UserService) CreateUser(name string, email string, password string) (models.User, error) {
	user := models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	result := db.Client.Create(&user)
	return user, result.Error
}

func (us *UserService) GetUserByID(ID uint) (models.User, error) {
	var user models.User
	result := db.Client.First(&user, ID)
	return user, result.Error
}

func (us *UserService) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	result := db.Client.Where("Email = ?", email).First(&user)
	return user, result.Error
}

func (us *UserService) GetUsersByTenantID(ID uint) ([]models.User, error) {
	var users []models.User
	result := db.Client.
		Preload("RoleAssignments", "TenantID = ?", ID).
		Where("RoleAssignments.TenantID = ?", ID).
		Find(&users)

	return users, result.Error
}
