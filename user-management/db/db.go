package db

import (
	"quickcooks/user-management/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Init() {
	dsn := "host=localhost user=quickcooks password=password dbname=quickcooks"
	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func Migrate() {
	Db.AutoMigrate(
		&models.Tenant{},
		&models.User{},
		&models.Permission{},
		&models.Role{},
		&models.RoleAssignment{},
		&models.RolePermission{},
	)
}
