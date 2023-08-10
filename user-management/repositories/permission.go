package repositories

import (
	"quickcooks/user-management/models"

	"gorm.io/gorm"
)

type IPermissionRepository interface {
	GetAll() ([]*models.Permission, error)
	GetByID(ID uint) (*models.Permission, error)
	Create(permission *models.Permission) (*models.Permission, error)
	Delete(permission *models.Permission) (*models.Permission, error)
}

type GormPermissionRepository struct {
	DB *gorm.DB
}

func NewGormPermissionRepository(db *gorm.DB) *GormPermissionRepository {
	return &GormPermissionRepository{
		DB: db,
	}
}

func (r *GormPermissionRepository) GetAll() ([]*models.Permission, error) {
	var permissions []*models.Permission
	result := r.DB.Find(&permissions)
	return permissions, result.Error
}

func (r *GormPermissionRepository) GetByID(ID uint) (*models.Permission, error) {
	var permission *models.Permission
	result := r.DB.Find(&permission, ID)
	return permission, result.Error
}

func (r *GormPermissionRepository) Create(permission *models.Permission) (*models.Permission, error) {
	result := r.DB.Create(&permission)
	return permission, result.Error
}

func (r *GormPermissionRepository) Delete(permission *models.Permission) (*models.Permission, error) {
	result := r.DB.Delete(&permission)
	return permission, result.Error
}
