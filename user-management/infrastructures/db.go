package infrastructures

import (
	"quickcooks/user-management/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDB() *gorm.DB {
	dsn := "host=localhost user=quickcooks password=password dbname=quickcooks"
	var err error
	client, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error connecting to database")
	}
	err = migrateDatabase(client, true)
	if err != nil {
		panic("Error migrating database")
	}
	return client
}

func dropAllTables(migrator gorm.Migrator) error {
	var err error
	err = migrator.DropTable(&models.Tenant{})
	if err != nil {
		return err
	}
	err = migrator.DropTable(&models.User{})
	if err != nil {
		return err
	}
	err = migrator.DropTable(&models.Permission{})
	if err != nil {
		return err
	}
	err = migrator.DropTable(&models.Role{})
	if err != nil {
		return err
	}
	err = migrator.DropTable(&models.RoleAssignment{})
	if err != nil {
		return err
	}
	err = migrator.DropTable(&models.RolePermission{})
	if err != nil {
		return err
	}

	return nil
}

func migrateDatabase(database *gorm.DB, drop bool) error {
	migrator := database.Migrator()

	if drop {
		err := dropAllTables(migrator)
		if err != nil {
			return err
		}
	}

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
