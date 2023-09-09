package models

type RoleAssignment struct {
	Entity
	TenantID uint
	UserID   uint
	RoleID   uint
}
