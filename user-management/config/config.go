package config

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	Environment      string
	ConnectionString string
	Seed             bool
	Serve            bool
	JwtSecret        string
}

func ReadConfig() *Config {
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

	jwtSecret, found := os.LookupEnv("JWT_SECRET")
	if !found {
		jwtSecret = "secret!"
	}

	return &Config{
		Environment:      *environment,
		ConnectionString: connectionString,
		Seed:             *seed,
		Serve:            *serve,
		JwtSecret:        jwtSecret,
	}
}
