package repositories

import (
	"quickcooks/user-management/models"

	"gorm.io/gorm"
)

type IRolePermissionRepository interface {
	GetByID(ID uint) (*models.RolePermission, error)
	GetByRoleID(roleID uint) ([]*models.RolePermission, error)
	Create(rolePermission *models.RolePermission) (*models.RolePermission, error)
	Delete(rolePermission *models.RolePermission) (*models.RolePermission, error)
}

type GormRolePermissionRepository struct {
	DB *gorm.DB
}

func NewGormRolePermissionRepository(db *gorm.DB) *GormRolePermissionRepository {
	return &GormRolePermissionRepository{
		DB: db,
	}
}

func (r *GormRolePermissionRepository) GetByID(ID uint) (*models.RolePermission, error) {
	var rolePermission *models.RolePermission
	result := r.DB.Find(&rolePermission, ID)
	return rolePermission, result.Error
}

func (r *GormRolePermissionRepository) GetByRoleID(roleID uint) ([]*models.RolePermission, error) {
	var rolePermissions []*models.RolePermission
	result := r.DB.
		Preload("Permission").
		Preload("Role").
		Where("RoleID = ?", roleID).
		Find(&rolePermissions)
	return rolePermissions, result.Error
}

func (r *GormRolePermissionRepository) Create(rolePermission *models.RolePermission) (*models.RolePermission, error) {
	result := r.DB.Create(rolePermission)
	return rolePermission, result.Error
}

func (r *GormRolePermissionRepository) Delete(rolePermission *models.RolePermission) (*models.RolePermission, error) {
	result := r.DB.Delete(rolePermission)
	return rolePermission, result.Error
}
