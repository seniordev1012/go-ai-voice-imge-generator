package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
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
	//CollectInstalledApps()
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

func CollectInstalledApps() {

	// Open an SQLite database
	db, err := sql.Open("sqlite3", "DB/installed_programs.db")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}(db)

	// Create a table to store program names
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS programs (name TEXT)")
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(err)
	// Run a Bash command to list installed programs and capture the output
	cmd := exec.Command("bash", "-c", "dpkg --get-selections | grep -v deinstall")
	log.Println(cmd)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	// Split the output into lines
	lines := strings.Split(string(output), "\n")
	log.Println(lines)

	// Prepare an SQL statement to insert program names
	stmt, err := db.Prepare("INSERT INTO programs (name) VALUES (?)")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(stmt)
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Println(err)
		}
	}(stmt)

	// Insert program names into the database
	for _, line := range lines {
		if line != "" {
			program := strings.Fields(line)[0]
			_, err := stmt.Exec(program)
			log.Println(program)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	fmt.Println("Installed programs have been stored in the SQLite database.")
}
