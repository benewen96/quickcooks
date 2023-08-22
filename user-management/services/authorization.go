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
) (*AuthorizationService, error) {
	authorizationService := &AuthorizationService{
		roleRepository:           roleRepository,
		rolePermissionRepository: rolePermissionRepository,
		permissionRepository:     permissionRepository,
	}

	roles, err := authorizationService.seedRoles()
	if err != nil {
		return nil, err
	}
	permissions, err := authorizationService.seedPermissions()
	if err != nil {
		return nil, err
	}

	err = authorizationService.seedRolePermissions(roles, permissions)
	if err != nil {
		return nil, err
	}

	return authorizationService, nil
}

// GetRoles returns a list of all application roles and associated permissions
func (s *AuthorizationService) GetRoles() ([]*models.Role, error) {
	return s.roleRepository.GetAll()
}

// GetRoleByName attempts to find a role with the given name if it exists
func (s *AuthorizationService) GetRoleByName(name string) (*models.Role, error) {
	return s.roleRepository.GetByName(name)
}

func (s *AuthorizationService) seedRoles() ([]*models.Role, error) {
	admin, err := s.roleRepository.FindOrCreate(&models.Role{Name: "admin"})
	if err != nil {
		return nil, err
	}

	member, err := s.roleRepository.FindOrCreate(&models.Role{Name: "member"})
	if err != nil {
		return nil, err
	}

	return []*models.Role{admin, member}, nil
}

func (s *AuthorizationService) seedPermissions() ([]*models.Permission, error) {
	var permissions []*models.Permission

	resources := []string{
		"recipe",
		"plan",
		"order",
		"food",
		"staple",
		"tenant",
	}

	actions := []string{
		"create",
		"read",
		"update",
		"delete",
	}

	for _, resource := range resources {
		for _, action := range actions {
			permission := &models.Permission{Resource: resource, Action: action}
			permission, err := s.permissionRepository.FindOrCreate(permission)

			if err != nil {
				return nil, err
			}

			permissions = append(permissions, permission)
		}
	}

	return permissions, nil
}
func (s *AuthorizationService) seedRolePermissions(roles []*models.Role, permissions []*models.Permission) error {
	for _, p := range permissions {
		_, err := s.rolePermissionRepository.
			FindOrCreate(&models.RolePermission{
				RoleID:       roles[0].ID,
				PermissionID: p.ID,
			})
		if err != nil {
			return err
		}

		if p.Action == "read" {
			_, err := s.rolePermissionRepository.FindOrCreate(&models.RolePermission{
				RoleID:       roles[1].ID,
				PermissionID: p.ID,
			})
			if err != nil {
				return err
			}
		}
	}
	return nil
}
