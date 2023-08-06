package models

import (
	"quickcooks/user-management/db"

	"gorm.io/gorm"
)

type Tenant struct {
	gorm.Model
	Name            string
	RoleAssignments []RoleAssignment
}

func (t *Tenant) UpdateName(name string) error {
	t.Name = name
	return db.Client.Save(&t).Error
}

func (t *Tenant) AssignRole(userID uint, roleID uint) (RoleAssignment, error) {
	roleAssignment := RoleAssignment{
		TenantID: t.ID,
		RoleID:   roleID,
		UserID:   userID,
	}
	result := db.Client.Create(&roleAssignment)
	return roleAssignment, result.Error
}
