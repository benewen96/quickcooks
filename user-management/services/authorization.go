package services

import (
	"quickcooks/user-management/models"
	"quickcooks/user-management/repositories"
)

type AuthorizationService struct {
	roleRepository           repositories.IRoleRepository
	rolePermissionRepository repositories.IRolePermissionRepository
	permissionRepository     repositories.IPermissionRepository
}

func NewAuthorizationService(
	roleRepository repositories.IRoleRepository,
	rolePermissionRepository repositories.IRolePermissionRepository,
	permissionRepository repositories.IPermissionRepository,
) *AuthorizationService {
	return &AuthorizationService{
		roleRepository:           roleRepository,
		rolePermissionRepository: rolePermissionRepository,
		permissionRepository:     permissionRepository,
	}
}

func (s *AuthorizationService) GetRoles() ([]*models.Role, error) {
	return s.roleRepository.GetAll()
}

func (s *AuthorizationService) GetRoleByName(name string) (*models.Role, error) {
	return s.roleRepository.GetByName(name)
}
