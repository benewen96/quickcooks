package services

import (
	"quickcooks/user-management/models"
	"quickcooks/user-management/repositories"
)

type MyTenantsService struct {
	TenantRepository         repositories.ITenantRepository
	RoleRepository           repositories.IRoleRepository
	RoleAssignmentRepository repositories.IRoleAssignmentRepository
}

func NewMyTenantsService(
	tenantRepository repositories.ITenantRepository,
	roleRepository repositories.IRoleRepository,
	roleAssignmentRepository repositories.IRoleAssignmentRepository,
) *MyTenantsService {
	return &MyTenantsService{
		TenantRepository:         tenantRepository,
		RoleRepository:           roleRepository,
		RoleAssignmentRepository: roleAssignmentRepository,
	}
}

func (s *MyTenantsService) GetTenantsByUserID(userID uint) ([]*models.Tenant, error) {
	return s.TenantRepository.GetByUserID(userID)
}

func (s *MyTenantsService) CreateTenantWithAdmin(name string, userID uint) (*models.Tenant, error) {
	tenant := &models.Tenant{
		Name: name,
	}
	tenant, err := s.TenantRepository.Create(tenant)
	if err != nil {
		return nil, err
	}

	role, err := s.RoleRepository.GetByName("admin")
	if err != nil {
		return nil, err
	}

	roleAssignment := &models.RoleAssignment{
		TenantID: tenant.ID,
		UserID:   userID,
		RoleID:   role.ID,
	}
	_, err = s.RoleAssignmentRepository.Create(roleAssignment)
	if err != nil {
		return nil, err
	}

	return s.TenantRepository.GetByID(tenant.ID)
}

func (s *MyTenantsService) UpdateTenantName(tenantID uint, name string) (*models.Tenant, error) {
	tenant, err := s.TenantRepository.GetByID(tenantID)
	if err != nil {
		return nil, err
	}
	return s.TenantRepository.UpdateName(tenant, name)
}

func (s *MyTenantsService) AssignTenantRole(tenantID uint, userID uint, roleID uint) (*models.RoleAssignment, error) {
	roleAssignment := &models.RoleAssignment{
		TenantID: tenantID,
		UserID:   userID,
		RoleID:   roleID,
	}
	return s.RoleAssignmentRepository.Create(roleAssignment)
}

func (s *MyTenantsService) UnassignTenantRole(roleAssignmentID uint) (*models.RoleAssignment, error) {
	roleAssignment, err := s.RoleAssignmentRepository.GetByID(roleAssignmentID)
	if err != nil {
		return nil, err
	}
	return s.RoleAssignmentRepository.Delete(roleAssignment)
}

func (s *MyTenantsService) UnassignTenantUser(userID uint) ([]*models.RoleAssignment, error) {
	roleAssignments, err := s.RoleAssignmentRepository.GetByUserID(userID)
	if err != nil {
		return nil, err
	}
	return s.RoleAssignmentRepository.DeleteMany(roleAssignments)
}
