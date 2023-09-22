package main

import (
	"fmt"
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
