package repositories

import (
	"quickcooks/user-management/models"

	"gorm.io/gorm"
)

// A IPermissionRepository provides management of permission data
type IPermissionRepository interface {
	GetAll() ([]*models.Permission, error)
	GetByID(ID uint) (*models.Permission, error)
	Create(permission *models.Permission) (*models.Permission, error)
	Delete(permission *models.Permission) (*models.Permission, error)
}

// A GormPermissionRepository provides management of permission data within the
// gorm database
type GormPermissionRepository struct {
	DB *gorm.DB
}

// NewGormPermissionRepository returns a new GormPermissionRepository instance
// with the given gorm database
func NewGormPermissionRepository(db *gorm.DB) *GormPermissionRepository {
	return &GormPermissionRepository{
		DB: db,
	}
}

// GetAll returns all permissions
func (r *GormPermissionRepository) GetAll() ([]*models.Permission, error) {
	var permissions []*models.Permission
	result := r.DB.Find(&permissions)
	return permissions, result.Error
}

// GetByID returns the permission witht the given ID, if it exists
func (r *GormPermissionRepository) GetByID(ID uint) (*models.Permission, error) {
	var permission *models.Permission
	result := r.DB.Find(&permission, ID)
	return permission, result.Error
}

// Create inserts the given permission
func (r *GormPermissionRepository) Create(permission *models.Permission) (*models.Permission, error) {
	result := r.DB.Create(&permission)
	return permission, result.Error
}

// Delete deletes the given permission
func (r *GormPermissionRepository) Delete(permission *models.Permission) (*models.Permission, error) {
	result := r.DB.Delete(&permission)
	return permission, result.Error
}
