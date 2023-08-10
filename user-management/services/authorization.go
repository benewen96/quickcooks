package services

import (
	"quickcooks/user-management/models"
	"quickcooks/user-management/repositories"
)

type AuthorizationService struct {
	RoleRepository           repositories.IRoleRepository
	RolePermissionRepository repositories.IRolePermissionRepository
	PermissionRepository     repositories.IPermissionRepository
}

func NewAuthorizationService(
	roleRepository repositories.IRoleRepository,
	rolePermissionRepository repositories.IRolePermissionRepository,
	permissionRepository repositories.IPermissionRepository,
) *AuthorizationService {
	return &AuthorizationService{
		RoleRepository:           roleRepository,
		RolePermissionRepository: rolePermissionRepository,
		PermissionRepository:     permissionRepository,
	}
}

func (s *AuthorizationService) seedRoles() ([]*models.Role, error) {
	var roles []*models.Role

	admin := &models.Role{Name: "admin"}
	admin, err := s.RoleRepository.Create(admin)
	if err != nil {
		return roles, err
	}

	member := &models.Role{Name: "member"}
	member, err = s.RoleRepository.Create(member)
	if err != nil {
		return roles, err
	}

	roles = []*models.Role{admin, member}

	return roles, err
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
			permission, err := s.PermissionRepository.Create(permission)

			if err != nil {
				return permissions, err
			}

			permissions = append(permissions, permission)
		}
	}

	return permissions, nil
}

func (s *AuthorizationService) GetRoles() ([]*models.Role, error) {
	return s.RoleRepository.GetAll()
}

func (s *AuthorizationService) GetRoleByName(name string) (*models.Role, error) {
	return s.RoleRepository.GetByName(name)
}

func (s *AuthorizationService) SeedAuthorizationData() error {
	roles, err := s.seedRoles()
	if err != nil {
		return err
	}
	permissions, err := s.seedPermissions()
	if err != nil {
		return err
	}

	for _, p := range permissions {
		rolePermission := &models.RolePermission{RoleID: roles[0].ID, PermissionID: p.ID}
		s.RolePermissionRepository.Create(rolePermission)
		if err != nil {
			return err
		}

		if p.Action == "read" {
			rolePermission = &models.RolePermission{RoleID: roles[1].ID, PermissionID: p.ID}
			s.RolePermissionRepository.Create(rolePermission)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
