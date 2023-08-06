package models

import (
	"quickcooks/user-management/db"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name            string
	Email           string
	Password        string
	RoleAssignments []RoleAssignment
}

func (u *User) UpdateName(name string) error {
	u.Name = name
	return db.Client.Save(&u).Error
}

func (u *User) UpdateEmail(email string) error {
	u.Email = email
	return db.Client.Save(&u).Error
}

func (u *User) UpdatePassword(password string) error {
	u.Password = password
	return db.Client.Save(&u).Error
}
