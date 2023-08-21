package repositories

import (
	"quickcooks/user-management/models"

	"gorm.io/gorm"
)

// A IRolePermissionRepository provides management of role permission data
type IRolePermissionRepository interface {
	GetByID(ID uint) (*models.RolePermission, error)
	GetByRoleID(roleID uint) ([]*models.RolePermission, error)
	Create(rolePermission *models.RolePermission) (*models.RolePermission, error)
	Delete(rolePermission *models.RolePermission) (*models.RolePermission, error)
}

// A GormRolePermissionRepository provides management of role permission data
// within the gorm database
type GormRolePermissionRepository struct {
	DB *gorm.DB
}

// NewGormRolePermissionRepository returns a new GormRolePermissionRepository
// instance with the given gorm database
func NewGormRolePermissionRepository(db *gorm.DB) *GormRolePermissionRepository {
	return &GormRolePermissionRepository{
		DB: db,
	}
}

// GetByID returns the role permission with the given ID
func (r *GormRolePermissionRepository) GetByID(ID uint) (*models.RolePermission, error) {
	var rolePermission *models.RolePermission
	result := r.DB.Find(&rolePermission, ID)
	return rolePermission, result.Error
}

// GetByRoleID returns all role permissions with the given role ID
func (r *GormRolePermissionRepository) GetByRoleID(roleID uint) ([]*models.RolePermission, error) {
	var rolePermissions []*models.RolePermission
	result := r.DB.
		Preload("Permission").
		Preload("Role").
		Where("RoleID = ?", roleID).
		Find(&rolePermissions)
	return rolePermissions, result.Error
}

// Create creates the given role permission
func (r *GormRolePermissionRepository) Create(rolePermission *models.RolePermission) (*models.RolePermission, error) {
	result := r.DB.Create(&rolePermission)
	return rolePermission, result.Error
}

// Delete deletes the given role permission
func (r *GormRolePermissionRepository) Delete(rolePermission *models.RolePermission) (*models.RolePermission, error) {
	result := r.DB.Delete(&rolePermission)
	return rolePermission, result.Error
}
