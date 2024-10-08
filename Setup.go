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

type dbFunc func() error

func executeDBFunc(fn dbFunc) {
	if err := fn(); err != nil {
		log.Printf("Error creating database: %v", err)
	}
}

func dbInit() error {
	executeDBFunc(extensionsSource)
	executeDBFunc(createMasterMessages)
	executeDBFunc(createGalleryDatabase)
	executeDBFunc(createMessagesDatabase)
	executeDBFunc(createUserDatabase)
	executeDBFunc(createSettingsDatabase)
	executeDBFunc(createLocalMediaDatabase)
	executeDBFunc(createProductivityDatabase)
	executeDBFunc(createKeyloggerDatabase)
	executeDBFunc(createLLMSelectionDatabase)
	executeDBFunc(createSpeechSelectionDatabase)
	return nil
}

func SetEnvironmentVariable(key string, tokenValue string) {
	valueToStore := os.Setenv(key, tokenValue)
	if valueToStore != nil {
		log.Printf("Error setting environment variable: %v", valueToStore)
	}
}

func _() {

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
