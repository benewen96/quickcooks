package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormClient() *gorm.DB {
	dsn := "host=localhost user=quickcooks password=password dbname=quickcooks"
	var err error
	client, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return client
}
