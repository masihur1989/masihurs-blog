package common

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/codingmechanics/applogger"

	"github.com/joho/godotenv"
)

// ConfigKeys -
var ConfigKeys = []string{
	"DB_HOST",
	"DB_USER",
	"DB_PASS",
	"DB_NAME",
	// auth
	"JWT_SECRET",
}

var l applogger.Logger

// ConfigureApp will load environment variables from  .env files and AWS Secrets Manager, panic'ing if any of them fail.
func ConfigureApp() {
	l.Info("Attempting to load config from .env file")

	if err := setupDotEnv(); err != nil {
		log.Panicf("Could not load .env file: %e", err)
	}

	missingRequiredKeys := []string{}

	for _, key := range ConfigKeys {
		if _, ok := os.LookupEnv(key); !ok {
			missingRequiredKeys = append(missingRequiredKeys, key)
		}
	}

	if len(missingRequiredKeys) > 0 {
		msg := fmt.Sprintf("App cannot start because it is missing the following environment variables: %s", strings.Join(missingRequiredKeys, ", "))
		log.Panicf(msg)
	}
}

// setupDotEnv will load config values from a .env file
func setupDotEnv() error {
	// Attempt to load it relatively (works for local dev)
	err := godotenv.Load()

	if err == nil {
		l.Info("Successfully loaded .env relatively. Finished loading config.")
		return nil
	}

	// Attempt to load it from the root of the docker container (works in the docker container on production)
	err = godotenv.Load("/.env")
	if err != nil {
		l.Errorf("Couldn't load a .env file. The app will rely on environment variables exclusively: ", err)
		return nil
	}

	l.Info("Successfully loaded /.env. Finished loading config.")
	return nil
}
