package infrastructures

import (
	"log"
	"os"
	"quickcooks/user-management/models"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGormDB(connString string) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,  // Slow SQL threshold
			LogLevel:                  logger.Error, // Log level
			IgnoreRecordNotFoundError: true,         // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,         // Don't include params in the SQL log
			Colorful:                  false,        // Disable color
		},
	)

	var err error
	client, err := gorm.Open(postgres.Open(connString), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("Error connecting to database: " + err.Error())
	}

	err = client.AutoMigrate(
		&models.Tenant{},
		&models.User{},
		&models.Permission{},
		&models.Role{},
		&models.RoleAssignment{},
		&models.RolePermission{},
	)
	if err != nil {
		panic("Error migrating database: " + err.Error())
	}

	return client
}
