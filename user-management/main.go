package main

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	environment      string
	seed             string
	connectionString string
}

func ReadConfig() *Config {
	flag.CommandLine.SetOutput(os.Stdout)
	environment := flag.String("environment", "development", "Environment to use [development/production]")
	seed := flag.String("seed", "none", "Seeds the QuickCooks database with data [none/required/dev]")

	flag.Parse()

	if *environment != "development" && *seed == "dev" {
		message := "Cannot seed development data a non-development environment!\n"
		panic(message)
	}

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
		seed:             *seed,
		connectionString: connectionString,
	}
}

func main() {
	config := ReadConfig()
	context := newUserManagementContext(*config)
	err := newRouter(context).Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		panic("Error starting router:\n\n" + err.Error())
	}
}
