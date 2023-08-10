package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Client *gorm.DB

func init() {
	dsn := "host=localhost user=quickcooks password=password dbname=quickcooks"
	var err error
	Client, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
