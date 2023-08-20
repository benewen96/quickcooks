package infrastructures

import (
	"quickcooks/user-management/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDB(connString string) *gorm.DB {
	var err error
	client, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		panic("Error connecting to database")
	}
	return client
}

func MigrateDatabase(database *gorm.DB) error {
	migrator := database.Migrator()

	err := migrator.AutoMigrate(
		&models.Tenant{},
		&models.User{},
		&models.Permission{},
		&models.Role{},
		&models.RoleAssignment{},
		&models.RolePermission{},
	)
	if err != nil {
		return err
	}

	return nil
}
