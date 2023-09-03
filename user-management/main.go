package main

import (
	"quickcooks/user-management/config"
	"quickcooks/user-management/context"
	"quickcooks/user-management/routes"
)

func main() {
	config := config.ReadConfig()

	context, err := context.NewUserManagementContext(config)
	if err != nil {
		panic("Error creating user management context:\n" + err.Error())
	}

	if config.Seed {
		if config.Environment != "development" {
			panic("Cannot seed development data in a non-devlopment environment!")
		}
		err := context.Seed()
		if err != nil {
			panic("Unable to seed development data:\n" + err.Error())
		}
	}

	if config.Serve {
		err = routes.NewRouter(context).Run() // listen and serve on 0.0.0.0:8080
		if err != nil {
			panic("Error starting router:\n" + err.Error())
		}
	}
}
