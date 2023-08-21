package repositories

import (
	"errors"
	"quickcooks/user-management/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetByID(ID uint) (*models.User, error)
	GetByTenantID(tenantID uint) ([]*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Exists(email string) bool
	Create(user *models.User) (*models.User, error)
	Delete(user *models.User) (*models.User, error)
	UpdateName(user *models.User, name string) (*models.User, error)
	UpdateEmail(user *models.User, email string) (*models.User, error)
	UpdatePassword(user *models.User, password string) (*models.User, error)
}

type GormUserRepository struct {
	DB *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{
		DB: db,
	}
}

func (r *GormUserRepository) GetByID(ID uint) (*models.User, error) {
	var user *models.User
	result := r.DB.First(&user, ID)
	return user, result.Error
}

func (r *GormUserRepository) GetByTenantID(tenantID uint) ([]*models.User, error) {
	var users []*models.User
	result := r.DB.
		Preload("RoleAssignments").
		Preload("RoleAssignments.Role").
		Preload("RoleAssignments.Role.RolePermissions").
		Where("RoleAssignments.TenantID = ?", tenantID).
		Find(&users)

	return users, result.Error
}

func (r *GormUserRepository) GetByEmail(email string) (*models.User, error) {
	user := models.User{}
	result := r.DB.Where("Email = ?", email).First(&user)
	return &user, result.Error
}

func (r *GormUserRepository) Exists(email string) bool {
	_, err := r.GetByEmail(email)
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

func (r *GormUserRepository) Create(user *models.User) (*models.User, error) {
	result := r.DB.Create(user)
	return user, result.Error
}

func (r *GormUserRepository) Delete(user *models.User) (*models.User, error) {
	result := r.DB.Delete(user)
	return user, result.Error
}

func (r *GormUserRepository) UpdateName(user *models.User, name string) (*models.User, error) {
	result := r.DB.First(user).Update("Name", name)
	return user, result.Error
}

func (r *GormUserRepository) UpdateEmail(user *models.User, email string) (*models.User, error) {
	user.Email = email
	result := r.DB.First(user).Update("Email", email)
	return user, result.Error
}

func (r *GormUserRepository) UpdatePassword(user *models.User, password string) (*models.User, error) {
	user.Password = password
	result := r.DB.First(user).Update("Password", password)
	return user, result.Error
}
