package main

import (
	"fmt"
	"quickcooks/user-management/db"
	"quickcooks/user-management/models"
	"quickcooks/user-management/server"
	"quickcooks/user-management/services"
)

type UserManagement struct {
	services.TenantService
	services.PermissionService
	services.RoleService
	services.UserService
}

var um UserManagement

func MigrateDatabase() {
	db.Client.AutoMigrate(
		&models.Tenant{},
		&models.User{},
		&models.Permission{},
		&models.Role{},
		&models.RoleAssignment{},
		&models.RolePermission{},
	)
}

func SeedDatabase() {
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

	admin, err := um.CreateRole("admin")
	if err != nil {
		fmt.Printf("Error creating role: %v", err)
	}

	member, err := um.CreateRole("member")
	if err != nil {
		fmt.Printf("Error creating role: %v", err)
	}

	for _, resource := range resources {
		for _, action := range actions {
			permission, err := um.CreatePermission(resource, action)

			if err != nil {
				fmt.Printf("Error creating permission: %v", err)
			}
			err = admin.AddPermission(permission)
			if err != nil {
				fmt.Printf("Error adding permission: %v", err)
			}
			if action == "read" {
				err = member.AddPermission(permission)
				if err != nil {
					fmt.Printf("Error adding permission: %v", err)
				}
			}
		}
	}

	joeBloggs, err := um.CreateUser("Joe Bloggs", "joe.bloggs@example.com", "password")
	if err != nil {
		fmt.Printf("Error creating user: %v", err)
	}

	janeBloggs, err := um.CreateUser("Jane Bloggs", "jane.bloggs@example.com", "password")
	if err != nil {
		fmt.Printf("Error creating user: %v", err)
	}

	tenant, err := um.CreateTenant("example_tenant")
	if err != nil {
		fmt.Printf("Error creating tenant: %v", err)
	}

	_, err = tenant.AssignRole(joeBloggs.ID, admin.ID)
	if err != nil {
		fmt.Printf("Error assigning role: %v", err)
	}
	_, err = tenant.AssignRole(janeBloggs.ID, member.ID)
	if err != nil {
		fmt.Printf("Error assigning role: %v", err)
	}
}

func main() {
	db.Init()
	MigrateDatabase()
	SeedDatabase()
	server.Init()
}
