package main

import (
	"quickcooks/user-management/db"
	"quickcooks/user-management/server"
)

func main() {
	db.Init()
	db.Migrate()
	server.Init()
}
