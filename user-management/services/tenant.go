package services

import (
	"quickcooks/user-management/db"
	"quickcooks/user-management/models"
)

type ITenantService interface {
	CreateTenant(name string) (models.Tenant, error)
	GetTenantByID(ID uint) (models.Tenant, error)
	GetTenantsByUserID(ID uint) ([]models.Tenant, error)
}

type TenantService struct{}

func (ts *TenantService) CreateTenant(name string) (models.Tenant, error) {
	tenant := models.Tenant{
		Name: name,
	}
	result := db.Client.Create(&tenant)
	return tenant, result.Error
}

func (ts *TenantService) GetTenantByID(ID uint) (models.Tenant, error) {
	var tenant models.Tenant
	result := db.Client.First(&tenant, ID)
	return tenant, result.Error
}

func (ts *TenantService) GetTenantsByUserID(ID uint) ([]models.Tenant, error) {
	var tenants []models.Tenant
	result := db.Client.
		Preload("RoleAssignments", "UserID = ?", ID).
		Where("RoleAssignments.UserID = ?", ID).
		Find(&tenants)

	return tenants, result.Error
}
