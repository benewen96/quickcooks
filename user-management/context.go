package main

import (
	"quickcooks/user-management/models"
	"quickcooks/user-management/repositories"
	"quickcooks/user-management/services"

	"gorm.io/gorm"
)

// An inversion of control container that registers all services for the user
// management context
type UserManagementContext struct {
	RegistrationService   *services.RegistrationService
	MyProfileService      *services.MyProfileService
	MyTenantsService      *services.MyTenantsService
	AuthorizationService  *services.AuthorizationService
	AuthenticationService *services.AuthenticationService
}

func newUserManagementContext(database *gorm.DB) (*UserManagementContext, error) {
	userRepository := repositories.NewGormUserRepository(database)
	tenantRepository := repositories.NewGormTenantRepository(database)
	roleRepository := repositories.NewGormRoleRepository(database)
	roleAssignmentRepository := repositories.NewGormRoleAssignmentRepository(database)
	rolePermissionRepository := repositories.NewGormRolePermissionRepository(database)
	permissionRepository := repositories.NewGormPermissionRepository(database)

	registrationService := services.NewRegistrationService(userRepository)
	myProfileService := services.NewMyProfileService(userRepository)
	myTenantsService := services.NewMyTenantsService(tenantRepository, roleRepository, roleAssignmentRepository)
	authorizationService, err := services.NewAuthorizationService(roleRepository, rolePermissionRepository, permissionRepository)
	if err != nil {
		return nil, err
	}
	authenticationService, err := services.NewAuthenticationService(userRepository)
	if err != nil {
		return nil, err
	}

	userManagementContext := &UserManagementContext{
		RegistrationService:   registrationService,
		MyProfileService:      myProfileService,
		MyTenantsService:      myTenantsService,
		AuthorizationService:  authorizationService,
		AuthenticationService: authenticationService,
	}

	return userManagementContext, nil
}

func (c *UserManagementContext) Seed() error {
	var joeBloggs, janeBloggs *models.User
	var bloggsTenant *models.Tenant
	var memberRole *models.Role

	var err error

	if !c.AuthenticationService.CheckUserEmailExists("joe.bloggs@example.com") {
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

	if !c.AuthenticationService.CheckUserEmailExists("jane.bloggs@example.com") {
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
