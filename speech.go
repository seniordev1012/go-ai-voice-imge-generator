package main

import (
	"aigen/aigenAudioAutoPlay"
	"aigen/aigenRest"
	"log"
)

var notificationSoundFile = "notification.mp3"

func pressPlayAudio(messageCall string) (bool, string) {
	//Azure Speech

	soundFileName, checkError := aigenRest.SpeakOut(messageCall)

	if checkError == nil {
		log.Println(soundFileName)
		playSound := aigenAudioAutoPlay.PlayAudioPlayback(soundFileName)
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
	err := aigenAudioAutoPlay.PlayAudioPlayback(notificationSoundFile)
	if err != nil {
		log.Println(err)
	}
}

func playVoiceNote(filename string) {
	err := aigenAudioAutoPlay.PlayAudioPlayback(filename)
	if err != nil {
		log.Println(err)
	}
}

func getAudioFile(message string) {
	audioFile, err := getAudio(message)
	if err != nil {
		log.Println(err)
	}
	playVoiceNote(audioFile)
}
