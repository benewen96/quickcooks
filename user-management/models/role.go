package models

type Role struct {
	Entity
	Name            string
	RoleAssignments []RoleAssignment
	RolePermissions []RolePermission
}
