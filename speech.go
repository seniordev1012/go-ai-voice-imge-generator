package main

import "log"

var notificationSoundFile = "notification.mp3"

func pressPlayAudio(messageCall string) (bool, string) {
	//Azure Speech

	soundFileName, checkError := speakOut(messageCall)

	if checkError == nil {
		log.Println(soundFileName)
		playSound := playAudioPlayback(soundFileName)
		if playSound != nil {
			return true, soundFileName
		}
		log.Printf("Sound file %s played successfully", soundFileName)
	} else {
		log.Println(checkError)
	}

	return true, soundFileName
}

func notificationSound() {
	err := playAudioPlayback(notificationSoundFile)
	if err != nil {
		log.Println(err)
	}
}
