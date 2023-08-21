package services

import (
	"quickcooks/user-management/models"
	"quickcooks/user-management/repositories"
)

// A RegistrationService is a provider for user authorization functionality
type AuthorizationService struct {
	roleRepository           repositories.IRoleRepository
	rolePermissionRepository repositories.IRolePermissionRepository
	permissionRepository     repositories.IPermissionRepository
}

// NewAuthorizationService creates a new AuthorizationService instance with the
// given role, permission, and rolePermission repositories
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

// GetRoles returns a list of all application roles and associated permissions
func (s *AuthorizationService) GetRoles() ([]*models.Role, error) {
	return s.roleRepository.GetAll()
}

// GetRoleByName attempts to find a role with the given name if it exists
func (s *AuthorizationService) GetRoleByName(name string) (*models.Role, error) {
	return s.roleRepository.GetByName(name)
}
