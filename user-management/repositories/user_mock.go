package repositories

import (
	"quickcooks/user-management/models"
)

type MockUserRepository struct {
	users []*models.User
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{}
}

func (r *MockUserRepository) GetByID(ID uint) (*models.User, error) {
	for _, u := range r.users {
		if u.ID == ID {
			return u, nil
		}
	}
	return nil, nil
}

func (r *MockUserRepository) GetByTenantID(tenantID uint) ([]*models.User, error) {
	var users []*models.User
	for _, u := range r.users {
		for _, ra := range u.RoleAssignments {
			if ra.TenantID == tenantID {
				users = append(users, u)
				break
			}
		}
	}
	return users, nil
}

func (r *MockUserRepository) GetByEmail(email string) (*models.User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, nil
}

func (r *MockUserRepository) Exists(email string) bool {
	user, _ := r.GetByEmail(email)
	return user == nil
}

func (r *MockUserRepository) Create(user *models.User) (*models.User, error) {
	r.users = append(r.users, user)
	return user, nil
}

func (r *MockUserRepository) Delete(user *models.User) (*models.User, error) {
	for i, u := range r.users {
		if u.ID != user.ID {
			continue
		}
		r.users = append(r.users[:i], r.users[i+1:]...)
	}
	return user, nil
}

func (r *MockUserRepository) UpdateName(user *models.User, name string) (*models.User, error) {
	for i, u := range r.users {
		if u.ID != user.ID {
			continue
		}
		r.users[i].Name = name
	}
	return user, nil
}

func (r *MockUserRepository) UpdateEmail(user *models.User, email string) (*models.User, error) {
	for i, u := range r.users {
		if u.ID != user.ID {
			continue
		}
		r.users[i].Email = email
	}
	return user, nil
}

func (r *MockUserRepository) UpdatePassword(user *models.User, password string) (*models.User, error) {
	for i, u := range r.users {
		if u.ID != user.ID {
			continue
		}
		r.users[i].Password = password
	}
	return user, nil
}
