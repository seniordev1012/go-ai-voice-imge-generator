package main

import (
	"fmt"
	"os/exec"
)

// LockSmith is used to set up the environment variables for the program to use
// By Running the azure.sh and slize.sh scripts
// See: https://learn.microsoft.com/en-us/azure/cognitive-services/speech-service/get-started-text-to-speech?tabs=linux%2Cterminal&pivots=programming-language-go#set-up-the-environment
func LockSmith() {
	// Run the first script
	cmd1 := exec.Command("bash", "azure.sh")
	if err := cmd1.Run(); err != nil {
		fmt.Println("Error running Azure Shell Script And Setting Up Your Keys:", err)
		return
	}

	// Run the second script
	cmd2 := exec.Command("bash", "slize.sh")
	if err := cmd2.Run(); err != nil {
		fmt.Println("Error running Slize Shell:", err)
		return
	}

	fmt.Println("Both scripts have finished running.")
}
