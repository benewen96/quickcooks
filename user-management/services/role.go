package services

import (
	"quickcooks/user-management/db"
	"quickcooks/user-management/models"
)

type IRoleService interface {
	CreateRole(name string) (models.Role, error)
	GetRoles() ([]models.Role, error)
}

type RoleService struct{}

func (rs *RoleService) CreateRole(name string) (models.Role, error) {
	role := models.Role{
		Name: name,
	}
	result := db.Client.Create(&role)
	return role, result.Error
}

func (rs *RoleService) GetRoles() ([]models.Role, error) {
	var roles []models.Role
	result := db.Client.Find(&roles)
	return roles, result.Error
}
