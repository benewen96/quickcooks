// # Models
//
// The models package provides the entity representations for use in object,
// relational mapping to the database
//
// The models within this subdomain are:
//   - User
//   - Tenant
//   - Permission
//   - Role
//   - Role Permission
//   - Role Assignment
package models

import (
	"time"

	"gorm.io/gorm"
)

type Entity struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
