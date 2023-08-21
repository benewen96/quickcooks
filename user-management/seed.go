package main

import (
	"quickcooks/user-management/models"
	"quickcooks/user-management/repositories"
)

type Seeder struct {
	roleRepository           repositories.IRoleRepository
	rolePermissionRepository repositories.IRolePermissionRepository
	permissionRepository     repositories.IPermissionRepository
	userRepository           repositories.IUserRepository
}

func NewSeeder(
	roleRepository repositories.IRoleRepository,
	rolePermissionRepository repositories.IRolePermissionRepository,
	permissionRepository repositories.IPermissionRepository,
	userRepository repositories.IUserRepository,
) *Seeder {
	return &Seeder{
		roleRepository:           roleRepository,
		rolePermissionRepository: rolePermissionRepository,
		permissionRepository:     permissionRepository,
		userRepository:           userRepository,
	}
}

func (s *Seeder) SeedRequiredData() error {
	roles, err := s.seedRoles()
	if err != nil {
		return err
	}
	permissions, err := s.seedPermissions()
	if err != nil {
		return err
	}

	return s.seedRolePermissions(roles, permissions)

}

func (s *Seeder) SeedDevData(c *UserManagementContext) error {
	var joeBloggs, janeBloggs *models.User
	var bloggsTenant *models.Tenant
	var memberRole *models.Role

	var err error

	if !s.userRepository.Exists("joe.bloggs@example.com") {
		joeBloggs, err = c.RegistrationService.RegisterUser("Joe Bloggs", "joe.bloggs@example.com", "password")
		if err != nil {
			return err
		}
		bloggsTenant, err = c.MyTenantsService.CreateTenantWithAdmin("Bloggs Tenant", joeBloggs.ID)
		if err != nil {
			return err
		}
	} else {
		joeBloggs, err = c.MyProfileService.GetUserByEmail("joe.bloggs@example.com")
		if err != nil {
			return err
		}
		tenants, err := c.MyTenantsService.GetTenantsByUserID(joeBloggs.ID)
		if err != nil {
			return err
		}
		if len(tenants) == 0 {
			bloggsTenant, err = c.MyTenantsService.CreateTenantWithAdmin("Bloggs Tenant", joeBloggs.ID)
			if err != nil {
				return err
			}
		} else {
			bloggsTenant, err = c.MyTenantsService.GetTenantByID(joeBloggs.RoleAssignments[0].TenantID)
			if err != nil {
				return err
			}
		}
	}

	if !s.userRepository.Exists("jane.bloggs@example.com") {
		janeBloggs, err = c.RegistrationService.RegisterUser("Jane Bloggs", "jane.bloggs@example.com", "password")
		if err != nil {
			return err
		}

		memberRole, err = c.AuthorizationService.GetRoleByName("member")
		if err != nil {
			return err
		}

		_, err = c.MyTenantsService.AssignTenantRole(bloggsTenant.ID, janeBloggs.ID, memberRole.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Seeder) seedRoles() ([]*models.Role, error) {
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

func (s *Seeder) seedPermissions() ([]*models.Permission, error) {
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
func (s *Seeder) seedRolePermissions(roles []*models.Role, permissions []*models.Permission) error {
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
