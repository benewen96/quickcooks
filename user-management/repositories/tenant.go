package repositories

import (
	"quickcooks/user-management/models"

	"gorm.io/gorm"
)

type ITenantRepository interface {
	GetByID(ID uint) (*models.Tenant, error)
	GetByUserID(userID uint) ([]*models.Tenant, error)
	Create(tenant *models.Tenant) (*models.Tenant, error)
	Delete(tenant *models.Tenant) (*models.Tenant, error)
	UpdateName(tenant *models.Tenant, name string) (*models.Tenant, error)
}

type GormTenantRepository struct {
	DB *gorm.DB
}

func NewGormTenantRepository(db *gorm.DB) *GormTenantRepository {
	return &GormTenantRepository{
		DB: db,
	}
}

func (r *GormTenantRepository) GetByID(ID uint) (*models.Tenant, error) {
	var tenant *models.Tenant
	result := r.DB.First(&tenant, ID)
	return tenant, result.Error
}

func (r *GormTenantRepository) GetByUserID(userID uint) ([]*models.Tenant, error) {
	var tenants []*models.Tenant
	result := r.DB.
		Preload("RoleAssignments").
		Preload("RoleAssignments.User").
		Preload("RoleAssignments.Role").
		Preload("RoleAssignments.Role.RolePermissions").
		Where("RoleAssignments.UserID = ?", userID).
		Find(&tenants)

	return tenants, result.Error
}

func (r *GormTenantRepository) Create(tenant *models.Tenant) (*models.Tenant, error) {
	result := r.DB.Create(&tenant)
	return tenant, result.Error
}

func (r *GormTenantRepository) Delete(tenant *models.Tenant) (*models.Tenant, error) {
	result := r.DB.Delete(&tenant)
	return tenant, result.Error
}

func (us *GormTenantRepository) UpdateName(tenant *models.Tenant, name string) (*models.Tenant, error) {
	result := us.DB.First(&tenant).Update("Name", name)
	return tenant, result.Error
}
