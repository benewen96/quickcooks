package main

import (
	"quickcooks/user-management/infrastructures"
	"quickcooks/user-management/repositories"
	"quickcooks/user-management/services"
)

type UserManagementContext struct {
	RegistrationService  *services.RegistrationService
	MyProfileService     *services.MyProfileService
	MyTenantsService     *services.MyTenantsService
	AuthorizationService *services.AuthorizationService
}

func newUserManagementContext() (*UserManagementContext, error) {
	var database = infrastructures.NewGormDB()

	var userRepository = repositories.NewGormUserRepository(database)
	var tenantRepository = repositories.NewGormTenantRepository(database)
	var roleRepository = repositories.NewGormRoleRepository(database)
	var roleAssignmentRepository = repositories.NewGormRoleAssignmentRepository(database)
	var rolePermissionRepository = repositories.NewGormRolePermissionRepository(database)
	var permissionRepository = repositories.NewGormPermissionRepository(database)

	var registrationService = services.NewRegistrationService(userRepository)
	var myProfileService = services.NewMyProfileService(userRepository)
	var myTenantsService = services.NewMyTenantsService(tenantRepository, roleRepository, roleAssignmentRepository)
	var authorizationService = services.NewAuthorizationService(roleRepository, rolePermissionRepository, permissionRepository)

	userManagementContext := &UserManagementContext{
		RegistrationService:  registrationService,
		MyProfileService:     myProfileService,
		MyTenantsService:     myTenantsService,
		AuthorizationService: authorizationService,
	}

	err := userManagementContext.seedDatabase()
	if err != nil {
		return nil, err
	}

	return userManagementContext, err
}

func (c *UserManagementContext) seedDatabase() error {
	err := c.AuthorizationService.SeedAuthorizationData()
	if err != nil {
		return err
	}

	joeBloggs, err := c.RegistrationService.RegisterUser("Joe Bloggs", "joe.bloggs@example.com", "password")
	if err != nil {
		return err
	}

	janeBloggs, err := c.RegistrationService.RegisterUser("Jane Bloggs", "jane.bloggs@example.com", "password")
	if err != nil {
		return err
	}

	tenant, err := c.MyTenantsService.CreateTenantWithAdmin("example_tenant", joeBloggs.ID)
	if err != nil {
		return err
	}

	member, err := c.AuthorizationService.GetRoleByName("member")
	if err != nil {
		return err
	}

	_, err = c.MyTenantsService.AssignTenantRole(tenant.ID, janeBloggs.ID, member.ID)
	if err != nil {
		return err
	}

	return nil
}
