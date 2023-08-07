package service

import (
	"quickcooks/user-management/db"
	"quickcooks/user-management/models"
)

type IPermissionService interface {
	CreatePermission(name string) (models.Permission, error)
	GetPermissions() ([]models.Permission, error)
}

type PermissionService struct{}

func (ps *PermissionService) CreatePermission(resource string, action string) (models.Permission, error) {
	permission := models.Permission{
		Resource: resource,
		Action:   action,
	}
	result := db.Client.Create(&permission)
	return permission, result.Error
}

func (ps *PermissionService) GetPermissions() ([]models.Permission, error) {
	var permissions []models.Permission
	result := db.Client.Find(&permissions)
	return permissions, result.Error
}
