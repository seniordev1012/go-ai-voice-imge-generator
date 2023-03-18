package main

import "log"

func pressPlayAudio(messageCall string) bool {
	soundFileName, checkError := speakOut(messageCall)
	if checkError == nil {
		log.Println(soundFileName)
		playSound := playAudio(soundFileName)
		if playSound != nil {
			return true
		}
		log.Printf("Sound file %s played successfully", soundFileName)
	} else {
		log.Println(checkError)
	}
	return false
}
