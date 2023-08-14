package main

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	environment  *string
	migrate      *bool
	seed         *bool
	pgConnString string
}

func ReadConfig() *Config {
	flag.CommandLine.SetOutput(os.Stdout)
	environment := flag.String("environment", "development", "Seed the QuickCooks database with initial data")
	migrate := flag.Bool("migrate", false, "Migrate the QuickCooks database")
	seed := flag.Bool("seed", false, "Seed the QuickCooks database with initial data")
	pgConnString, found := os.LookupEnv("PG_CONNECTION_STRING")
	flag.Parse()
	// shit code fix
	if !found && *environment == "development" {
		pgConnString = "host=localhost user=quickcooks password=password dbname=quickcooks"
	}
	return &Config{
		environment:  environment,
		migrate:      migrate,
		seed:         seed,
		pgConnString: pgConnString,
	}
}

func main() {
	config := ReadConfig()
	context := newUserManagementContext(*config)
	err := newRouter(context).Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		fmt.Printf("Error starting router: %v", err)
		return
	}
}
