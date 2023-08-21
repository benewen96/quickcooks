package main

import (
	"quickcooks/user-management/infrastructures"
	"quickcooks/user-management/repositories"
	"quickcooks/user-management/services"
)

// An inversion of control container that registers all services for the user
// management context
type UserManagementContext struct {
	RegistrationService  *services.RegistrationService
	MyProfileService     *services.MyProfileService
	MyTenantsService     *services.MyTenantsService
	AuthorizationService *services.AuthorizationService
}

func newUserManagementContext(config Config) *UserManagementContext {
	var database = infrastructures.NewGormDB(config.connectionString)

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

	if config.seed == "none" {
		return userManagementContext
	}

	seeder := NewSeeder(roleRepository, rolePermissionRepository, permissionRepository, userRepository)

	err := seeder.SeedRequiredData()
	if err != nil {
		panic("Unable to seed required data: " + err.Error())
	}

	if config.seed == "dev" {
		if config.environment != "development" {
			panic("Cannot seed development data in a non-devlopment environment!")
		}
		err := seeder.SeedDevData(userManagementContext)
		if err != nil {
			panic("Unable to seed development data: " + err.Error())
		}
	}

	return userManagementContext
}
