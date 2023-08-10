package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name            string
	RoleAssignments []RoleAssignment
	RolePermissions []RolePermission
}
