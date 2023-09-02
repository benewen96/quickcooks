package services

import (
	"quickcooks/user-management/models"
	"quickcooks/user-management/repositories"
)

// A MyTenantService is a provider for tenant functionality
type MyTenantsService struct {
	tenantRepository         repositories.ITenantRepository
	roleRepository           repositories.IRoleRepository
	roleAssignmentRepository repositories.IRoleAssignmentRepository
}

// NewMyTenantService creates a new MyTenantService instance with the given
// tenant, role, and roleAssignment repositories
func NewMyTenantsService(
	tenantRepository repositories.ITenantRepository,
	roleRepository repositories.IRoleRepository,
	roleAssignmentRepository repositories.IRoleAssignmentRepository,
) *MyTenantsService {
	return &MyTenantsService{
		tenantRepository:         tenantRepository,
		roleRepository:           roleRepository,
		roleAssignmentRepository: roleAssignmentRepository,
	}
}

// GetTenantByID returns the tenant with the given ID, if it exists
func (s *MyTenantsService) GetTenantByID(ID uint) (*models.Tenant, error) {
	return s.tenantRepository.GetByID(ID)
}

// GetTenantsByUserID returns all tenants where the user with the given ID has a role
// assignment
func (s *MyTenantsService) GetTenantsByUserID(userID uint) ([]*models.Tenant, error) {
	return s.tenantRepository.GetByUserID(userID)
}

// CreateTenantWithAdmin creates a new tenant with the given name and adds an
// admin role assignment to the user with the given ID
func (s *MyTenantsService) CreateTenantWithAdmin(name string, userID uint) (*models.Tenant, error) {
	tenant := &models.Tenant{
		Name: name,
	}
	tenant, err := s.tenantRepository.Create(tenant)
	if err != nil {
		return nil, err
	}

	role, err := s.roleRepository.GetByName("admin")
	if err != nil {
		return nil, err
	}

	roleAssignment := &models.RoleAssignment{
		TenantID: tenant.ID,
		UserID:   userID,
		RoleID:   role.ID,
	}
	_, err = s.roleAssignmentRepository.Create(roleAssignment)
	if err != nil {
		return nil, err
	}

	return s.tenantRepository.GetByID(tenant.ID)
}

// UpdateTenatName updates name of the tenant with the given ID, if it exists,
// to the given name
func (s *MyTenantsService) UpdateTenantName(tenantID uint, name string) (*models.Tenant, error) {
	tenant, err := s.tenantRepository.GetByID(tenantID)
	if err != nil {
		return nil, err
	}
	return s.tenantRepository.UpdateName(tenant, name)
}

// AssignTenantRole creates a new role assignment with the given IDs, given
// each target resource exists
func (s *MyTenantsService) AssignTenantRole(tenantID uint, userID uint, roleID uint) (*models.RoleAssignment, error) {
	roleAssignment := &models.RoleAssignment{
		TenantID: tenantID,
		UserID:   userID,
		RoleID:   roleID,
	}
	return s.roleAssignmentRepository.Create(roleAssignment)
}

// UnassignTenantRole removes the role assignment with the given ID, if it
// exists
func (s *MyTenantsService) UnassignTenantRole(roleAssignmentID uint) (*models.RoleAssignment, error) {
	roleAssignment, err := s.roleAssignmentRepository.GetByID(roleAssignmentID)
	if err != nil {
		return nil, err
	}
	return s.roleAssignmentRepository.Delete(roleAssignment)
}

// UnassignTenantUser removes all role assignments with the given user ID
func (s *MyTenantsService) UnassignTenantUser(userID uint) ([]*models.RoleAssignment, error) {
	roleAssignments, err := s.roleAssignmentRepository.GetByUserID(userID)
	if err != nil {
		return nil, err
	}
	return s.roleAssignmentRepository.DeleteMany(roleAssignments)
}
