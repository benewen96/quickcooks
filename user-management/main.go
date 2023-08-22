package main

import (
	"flag"
	"fmt"
	"os"
	"quickcooks/user-management/infrastructures"
)

type Config struct {
	environment      string
	connectionString string
	seed             bool
	serve            bool
}

func readConfig() *Config {
	flag.CommandLine.SetOutput(os.Stdout)
	environment := flag.String("environment", "development", "Environment to use [development/production]")
	seed := flag.Bool("seed", false, "Seeds the QuickCooks database with dummy data for development")
	serve := flag.Bool("serve", false, "Starts the application server")

	flag.Parse()

	connectionString, found := os.LookupEnv("PG_CONNECTION_STRING")
	if !found {
		if *environment == "development" {
			connectionString = "host=localhost user=quickcooks password=password dbname=quickcooks"
			message := "Environment variable 'PG_CONNECTION_STRING' is not set, defaulting to local environment:\n"
			message += "\t\"" + connectionString + "\"\n"
			fmt.Println(message)
		} else {
			message := "Environment variable 'PG_CONNECTION_STRING' is not set, cannot default in production\n\n"
			message += "Change environment or run the following, substituting where necessary for the target postgres database:\n"
			message += "\texport PG_CONNECTION_STRING=\"host=<hostname> user=<username> password=<password> dbname=<dbname>\"\n"
			panic(message)
		}
	}

	return &Config{
		environment:      *environment,
		connectionString: connectionString,
		seed:             *seed,
		serve:            *serve,
	}
}

func main() {
	config := readConfig()

	database := infrastructures.NewGormDB(config.connectionString)
	err := database.Error
	if err != nil {
		panic("Error connecting to database:\n" + err.Error())
	}

	context, err := newUserManagementContext(database)
	if err != nil {
		panic("Error creating user management context:\n" + err.Error())
	}

	if config.seed {
		if config.environment != "development" {
			panic("Cannot seed development data in a non-devlopment environment!")
		}
		err := context.Seed()
		if err != nil {
			panic("Unable to seed development data:\n" + err.Error())
		}
	}

	if config.serve {
		err = newRouter(context).Run() // listen and serve on 0.0.0.0:8080
		if err != nil {
			panic("Error starting router:\n" + err.Error())
		}
	}
}
