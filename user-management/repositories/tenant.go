package repositories

import (
	"quickcooks/user-management/models"

	"gorm.io/gorm"
)

// A ITenantRepository provides management of tenant data
type ITenantRepository interface {
	GetByID(ID uint) (*models.Tenant, error)
	GetByUserID(userID uint) ([]*models.Tenant, error)
	Create(tenant *models.Tenant) (*models.Tenant, error)
	Delete(tenant *models.Tenant) (*models.Tenant, error)
	UpdateName(tenant *models.Tenant, name string) (*models.Tenant, error)
}

// A GormTenantRepository provides management of tenant data within the gorm
// database
type GormTenantRepository struct {
	DB *gorm.DB
}

// NewGormTenantRepository returns a new GormTenantRepository instance with the
// given gorm database
func NewGormTenantRepository(db *gorm.DB) *GormTenantRepository {
	return &GormTenantRepository{
		DB: db,
	}
}

// GetByID returns the tenant with the given ID, if it exists
func (r *GormTenantRepository) GetByID(ID uint) (*models.Tenant, error) {
	var tenant *models.Tenant
	result := r.DB.First(&tenant, ID)
	return tenant, result.Error
}

// GetByUserID returns all tenants where the user with the given ID has a role
// assignment
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

// Create creates the given tenant
func (r *GormTenantRepository) Create(tenant *models.Tenant) (*models.Tenant, error) {
	result := r.DB.Create(&tenant)
	return tenant, result.Error
}

// Delete deletes the given tenant
func (r *GormTenantRepository) Delete(tenant *models.Tenant) (*models.Tenant, error) {
	result := r.DB.Delete(&tenant)
	return tenant, result.Error
}

// UpdateName updates the name of the tenant with the given ID, to the given name
func (us *GormTenantRepository) UpdateName(tenant *models.Tenant, name string) (*models.Tenant, error) {
	result := us.DB.First(&tenant).Update("Name", name)
	return tenant, result.Error
}
