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
	LockSmith()
	dbUser := os.Setenv("DB_USER", "")
	if dbUser != nil {
		log.Printf("Error setting environment variable: %v", dbUser)
	}
	dbPassed := os.Setenv("DB_PASSWORD", "")
	if dbPassed != nil {
		log.Printf("Error setting environment variable: %v", dbPassed)
	}
	dbHost := os.Setenv("DB_HOST", "")
	if dbHost != nil {
		log.Printf("Error setting environment variable: %v", dbHost)
	}
	dbName := os.Setenv("DB_NAME", "")
	if dbName != nil {
		log.Printf("Error setting environment variable: %v", dbName)
	}
	azureSpeechKey := os.Setenv("SPEECH_KEY", speechKeys)
	if azureSpeechKey != nil {
		log.Printf("Error setting environment variable: %v", azureSpeechKey)
	}

	openAiApiKey := os.Setenv("OPENAI", openKeys)
	if openAiApiKey != nil {
		log.Printf("Error setting environment variable: %v", openAiApiKey)
	}

}

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
