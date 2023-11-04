package main

import (
	"aigen/aigenRest"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
)

const (
	SetupPlatform           = "slize.sh"
	ExportAzureSpeechValues = "azure.sh"
)

// SetupScripts is used to set up the environment variables for the program to use
// By Running the azure.sh and slize.sh scripts
// See: https://learn.microsoft.com/en-us/azure/cognitive-services/speech-service/get-started-text-to-speech
// ?tabs=linux%2Cterminal&pivots=programming-language-go#set-up-the-environment
func SetupScripts() {
	RunBashScript(ExportAzureSpeechValues)
	RunBashScript(SetupPlatform)
	//Record if the setup scripts have finished running and don't run them again
	fmt.Println("Setup scripts have finished running.")
}

// RunBashScript will execute any shell script command given location of file
// Default is Platform root directory
func RunBashScript(shellScript string) bool {
	cmd2 := exec.Command("bash", shellScript)
	if err := cmd2.Run(); err != nil {
		fmt.Println("Error running Bash Script:", err)
		return true
	}
	return false
}

// RunProgram
func RunProgram(command string) bool {
	const PROGRAMSDB = "DB/installed_programs.db"
	db, err := sql.Open("sqlite3", PROGRAMSDB)
	if err != nil {
		fmt.Printf("Error opening the database: %v\n", err)
		return false
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}(db)

	log.Println(command, "Passed command")
	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM programs WHERE name = ?)", command).Scan(&exists)
	if err != nil {
		fmt.Printf("Error checking if the program exists: %v\n", err)
		return false
	}
	if !exists {
		fmt.Printf("Program '%s' is not installed.\n", command)
		return false
	}
	aigenRest.SendNotificationNow(fmt.Sprintf("Opening %s", command))
	// Execute the program
	cmd := exec.Command(command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Error running Bash Script:", err)
		return true
	}
	return false
}
