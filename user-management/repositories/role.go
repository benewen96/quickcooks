package repositories

import (
	"errors"
	"quickcooks/user-management/models"

	"gorm.io/gorm"
)

type IRoleRepository interface {
	GetAll() ([]*models.Role, error)
	GetByID(ID uint) (*models.Role, error)
	GetByName(name string) (*models.Role, error)
	Exists(name string) bool
	Create(role *models.Role) (*models.Role, error)
	FindOrCreate(role *models.Role) (*models.Role, error)
	Delete(role *models.Role) (*models.Role, error)
	UpdateName(role *models.Role, name string) (*models.Role, error)
}

type GormRoleRepository struct {
	DB *gorm.DB
}

func NewGormRoleRepository(db *gorm.DB) *GormRoleRepository {
	return &GormRoleRepository{
		DB: db,
	}
}

func (r *GormRoleRepository) GetAll() ([]*models.Role, error) {
	var roles []*models.Role
	result := r.DB.Find(&roles)
	return roles, result.Error
}

func (r *GormRoleRepository) GetByID(ID uint) (*models.Role, error) {
	var role *models.Role
	result := r.DB.Find(&role, ID)
	return role, result.Error
}

func (r *GormRoleRepository) GetByName(name string) (*models.Role, error) {
	var role *models.Role
	result := r.DB.Where("Name = ?", name).Find(&role)
	return role, result.Error
}

func (r *GormRoleRepository) Exists(name string) bool {
	_, err := r.GetByName(name)
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

// Create is not idempotent
func (r *GormRoleRepository) Create(role *models.Role) (*models.Role, error) {
	result := r.DB.Create(role)
	return role, result.Error
}

// FindOrCreate is idempotent
func (r *GormRoleRepository) FindOrCreate(role *models.Role) (*models.Role, error) {
	result := r.DB.FirstOrCreate(role, role)
	return role, result.Error
}

func (r *GormRoleRepository) Delete(role *models.Role) (*models.Role, error) {
	result := r.DB.Delete(role)
	return role, result.Error
}

func (r *GormRoleRepository) UpdateName(role *models.Role, name string) (*models.Role, error) {
	result := r.DB.First(role).Update("Name", name)
	return role, result.Error
}
