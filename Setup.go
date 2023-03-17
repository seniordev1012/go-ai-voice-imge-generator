package main

import (
	"log"
	"os"
)

func setup() {
	err := dbInit()
	if err != nil {
		log.Printf("Error initializing database: %v", err)
	}
	MigrationAssist()
}

func MigrationAssist() {
	err := os.Setenv("DB_USER", "")
	if err != nil {
		log.Printf("Error setting environment variable: %v", err)
	}
	err = os.Setenv("DB_PASSWORD", "")
	if err != nil {
		log.Printf("Error setting environment variable: %v", err)
	}
	err = os.Setenv("DB_HOST", "")
	if err != nil {
		log.Printf("Error setting environment variable: %v", err)
	}
	err = os.Setenv("DB_NAME", "")
	if err != nil {
		log.Printf("Error setting environment variable: %v", err)
	}

	err = os.Setenv("OPENAI", "")
	if err != nil {
		log.Printf("Error setting environment variable: %v", err)
	}
}

func dbInit() error {
	majorKeys := createKeyloggerDatabase()
	if majorKeys != nil {
		log.Printf("Error creating database: %v", majorKeys)
	}
	return error(nil)
}
