package repositories

import (
	"quickcooks/user-management/models"

	"gorm.io/gorm"
)

type IRoleAssignmentRepository interface {
	GetByID(ID uint) (*models.RoleAssignment, error)
	GetByTenantID(tenantID uint) ([]*models.RoleAssignment, error)
	GetByUserID(userID uint) ([]*models.RoleAssignment, error)
	Create(roleAssignment *models.RoleAssignment) (*models.RoleAssignment, error)
	Delete(roleAssignment *models.RoleAssignment) (*models.RoleAssignment, error)
	DeleteMany(roleAssignments []*models.RoleAssignment) ([]*models.RoleAssignment, error)
}

type GormRoleAssignmentRepository struct {
	DB *gorm.DB
}

func NewGormRoleAssignmentRepository(db *gorm.DB) *GormRoleAssignmentRepository {
	return &GormRoleAssignmentRepository{
		DB: db,
	}
}

func (r *GormRoleAssignmentRepository) GetByID(ID uint) (*models.RoleAssignment, error) {
	var roleAssignment *models.RoleAssignment
	result := r.DB.Find(&roleAssignment, ID)
	return roleAssignment, result.Error
}

func (r *GormRoleAssignmentRepository) GetByTenantID(tenantID uint) ([]*models.RoleAssignment, error) {
	var roleAssignments []*models.RoleAssignment
	result := r.DB.
		Preload("User").
		Preload("Role").
		Preload("Tenant").
		Where("TenantID = ?", tenantID).
		Find(&roleAssignments)
	return roleAssignments, result.Error
}

func (r *GormRoleAssignmentRepository) GetByUserID(userID uint) ([]*models.RoleAssignment, error) {
	var roleAssignments []*models.RoleAssignment
	result := r.DB.
		Preload("User").
		Preload("Role").
		Preload("Tenant").
		Where("UserID = ?", userID).
		Find(&roleAssignments)
	return roleAssignments, result.Error
}

func (r *GormRoleAssignmentRepository) Create(roleAssignment *models.RoleAssignment) (*models.RoleAssignment, error) {
	result := r.DB.Create(roleAssignment)
	return roleAssignment, result.Error
}

func (r *GormRoleAssignmentRepository) Delete(roleAssignment *models.RoleAssignment) (*models.RoleAssignment, error) {
	result := r.DB.Delete(roleAssignment)
	return roleAssignment, result.Error
}

func (r *GormRoleAssignmentRepository) DeleteMany(roleAssignments []*models.RoleAssignment) ([]*models.RoleAssignment, error) {
	tx := r.DB.Begin()
	for _, ra := range roleAssignments {
		err := tx.Delete(ra).Error
		if err != nil {
			tx = tx.Rollback()
			break
		}
	}
	tx = tx.Commit()
	return roleAssignments, tx.Error
}
