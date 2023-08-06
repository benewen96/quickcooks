package models

import (
	"quickcooks/user-management/db"

	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name            string
	RoleAssignments []RoleAssignment
	RolePermissions []RolePermission
}

func (r *Role) AddPermission(permission Permission) error {
	rolePermission := RolePermission{
		RoleID:       r.ID,
		PermissionID: permission.ID,
	}
	return db.Client.Create(&rolePermission).Error
}
