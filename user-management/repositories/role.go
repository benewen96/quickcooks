package repositories

import (
	"quickcooks/user-management/models"

	"gorm.io/gorm"
)

// A IRoleRepository provides management of role data
type IRoleRepository interface {
	GetAll() ([]*models.Role, error)
	GetByID(ID uint) (*models.Role, error)
	GetByName(name string) (*models.Role, error)
	Create(role *models.Role) (*models.Role, error)
	Delete(role *models.Role) (*models.Role, error)
	UpdateName(role *models.Role, name string) (*models.Role, error)
}

// A GormRoleRepository provides management of role data within the gorm
// database
type GormRoleRepository struct {
	DB *gorm.DB
}

// NewGormRoleRepository returns a new GormRoleRepository instance with the
// given gorm database
func NewGormRoleRepository(db *gorm.DB) *GormRoleRepository {
	return &GormRoleRepository{
		DB: db,
	}
}

// GetAll returns all roles
func (r *GormRoleRepository) GetAll() ([]*models.Role, error) {
	var roles []*models.Role
	result := r.DB.Find(&roles)
	return roles, result.Error
}

// GetByID returns the role with the given ID, if it exists
func (r *GormRoleRepository) GetByID(ID uint) (*models.Role, error) {
	var role *models.Role
	result := r.DB.Find(&role, ID)
	return role, result.Error
}

// GetByName returns the role with the given name, if it exists
func (r *GormRoleRepository) GetByName(name string) (*models.Role, error) {
	var role *models.Role
	result := r.DB.Where("Name = ?", name).Find(&role)
	return role, result.Error
}

// Create creates the role
func (r *GormRoleRepository) Create(role *models.Role) (*models.Role, error) {
	result := r.DB.Create(&role)
	return role, result.Error
}

// Delete deletes the role
func (r *GormRoleRepository) Delete(role *models.Role) (*models.Role, error) {
	result := r.DB.Delete(&role)
	return role, result.Error
}

// UpdateName updates the name of the role with the given ID, to the given name
func (r *GormRoleRepository) UpdateName(role *models.Role, name string) (*models.Role, error) {
	// TODO: Does this need to exist?
	result := r.DB.First(&role).Update("Name", name)
	return role, result.Error
}
