package models

type Tenant struct {
	Entity
	Name            string
	RoleAssignments []RoleAssignment
}
