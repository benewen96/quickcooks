package models

type Permission struct {
	Entity
	Resource        string
	Action          string
	RolePermissions []RolePermission
}
