package models

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Resource        string
	Action          string
	RolePermissions []RolePermission
}
