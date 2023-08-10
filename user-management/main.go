package main

import (
	"fmt"
	"quickcooks/user-management/infrastructures"
	"quickcooks/user-management/models"
	"quickcooks/user-management/repositories"
	"quickcooks/user-management/services"

	"gorm.io/gorm"
)

type UserManagementContext struct {
	RegistrationService  *services.RegistrationService
	MyProfileService     *services.MyProfileService
	MyTenantsService     *services.MyTenantsService
	AuthorizationService *services.AuthorizationService
}

func newUserManagementContext() (*UserManagementContext, error) {
	var database = infrastructures.NewGormDB()

	migrator := database.Migrator()

	err := dropAllTables(migrator)
	if err != nil {
		fmt.Printf("Error dropping tables: %v", err)
	}

	err = migrateDatabase(database, true)
	if err != nil {
		fmt.Printf("Error migrating database: %v", err)
	}

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

	err = userManagementContext.seedDatabase()
	if err != nil {
		return nil, err
	}

	return userManagementContext, err
}

func dropAllTables(migrator gorm.Migrator) error {
	var err error
	err = migrator.DropTable(&models.Tenant{})
	if err != nil {
		return err
	}
	err = migrator.DropTable(&models.User{})
	if err != nil {
		return err
	}
	err = migrator.DropTable(&models.Permission{})
	if err != nil {
		return err
	}
	err = migrator.DropTable(&models.Role{})
	if err != nil {
		return err
	}
	err = migrator.DropTable(&models.RoleAssignment{})
	if err != nil {
		return err
	}
	err = migrator.DropTable(&models.RolePermission{})
	if err != nil {
		return err
	}

	return nil
}

func migrateDatabase(database *gorm.DB, drop bool) error {
	migrator := database.Migrator()

	if drop {
		err := dropAllTables(migrator)
		if err != nil {
			return err
		}
	}

	err := migrator.AutoMigrate(
		&models.Tenant{},
		&models.User{},
		&models.Permission{},
		&models.Role{},
		&models.RoleAssignment{},
		&models.RolePermission{},
	)
	if err != nil {
		return err
	}

	return nil
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

func main() {
	context, err := newUserManagementContext()
	if err != nil {
		fmt.Printf("Error creating context: %v", err)
	}
	err = newRouter(context).Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		fmt.Printf("Error starting router: %v", err)
	}
}
