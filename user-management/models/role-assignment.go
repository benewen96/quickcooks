package models

import "gorm.io/gorm"

type RoleAssignment struct {
	gorm.Model
	TenantID uint
	UserID   uint
	RoleID   uint
}
