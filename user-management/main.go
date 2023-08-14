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
	devSeed      *bool
	pgConnString string
}

func ReadConfig() *Config {
	flag.CommandLine.SetOutput(os.Stdout)
	environment := flag.String("environment", "development", "Environment to use [development/production]")
	migrate := flag.Bool("migrate", false, "Migrate the QuickCooks database")
	seed := flag.Bool("seed", false, "Seed the QuickCooks database with required data")
	devSeed := flag.Bool("devSeed", false, "Seed the QuickCooks database with development data")
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
		devSeed:      devSeed,
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
