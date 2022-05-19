package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvConfig(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error cannot read .env file: %s", err.Error())
	}

	return os.Getenv(key)
}
