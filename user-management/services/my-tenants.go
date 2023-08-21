package services

import (
	"quickcooks/user-management/models"
	"quickcooks/user-management/repositories"
)

type MyTenantsService struct {
	tenantRepository         repositories.ITenantRepository
	roleRepository           repositories.IRoleRepository
	roleAssignmentRepository repositories.IRoleAssignmentRepository
}

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

func (s *MyTenantsService) GetTenantByID(ID uint) (*models.Tenant, error) {
	return s.tenantRepository.GetByID(ID)
}

func (s *MyTenantsService) GetTenantsByUserID(userID uint) ([]*models.Tenant, error) {
	return s.tenantRepository.GetByUserID(userID)
}

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

func (s *MyTenantsService) UpdateTenantName(tenantID uint, name string) (*models.Tenant, error) {
	tenant, err := s.tenantRepository.GetByID(tenantID)
	if err != nil {
		return nil, err
	}
	return s.tenantRepository.UpdateName(tenant, name)
}

func (s *MyTenantsService) AssignTenantRole(tenantID uint, userID uint, roleID uint) (*models.RoleAssignment, error) {
	roleAssignment := &models.RoleAssignment{
		TenantID: tenantID,
		UserID:   userID,
		RoleID:   roleID,
	}
	return s.roleAssignmentRepository.Create(roleAssignment)
}

func (s *MyTenantsService) UnassignTenantRole(roleAssignmentID uint) (*models.RoleAssignment, error) {
	roleAssignment, err := s.roleAssignmentRepository.GetByID(roleAssignmentID)
	if err != nil {
		return nil, err
	}
	return s.roleAssignmentRepository.Delete(roleAssignment)
}

func (s *MyTenantsService) UnassignTenantUser(userID uint) ([]*models.RoleAssignment, error) {
	roleAssignments, err := s.roleAssignmentRepository.GetByUserID(userID)
	if err != nil {
		return nil, err
	}
	return s.roleAssignmentRepository.DeleteMany(roleAssignments)
}
