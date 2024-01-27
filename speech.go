package main

import (
	"aigen/aigenAudioAutoPlay"
	"aigen/aigenRest"
	"log"
)

const notificationSoundFile = "notification.mp3"
const welcomeNotificationSoundFile = "welcome.mp3"

func pressPlayAudio(messageCall string) (bool, string) {
	//Azure Speech
	//TODO:: Allow switching of speech providers to allow for easy management of resources in cases
	//were Azure Speech starts to act real funny we can pull the plug real quick
	//soundFileName, checkError := aigenRest.SpeakOut(messageCall)
	//soundFileName, checkError := aigenRest.GCloudSpeakOUt(messageCall)
	soundFileName, checkError := aigenRest.GptSpeakOut(messageCall)

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

func playWelcomeSound() {
	err := aigenAudioAutoPlay.PlayAudioPlayback(welcomeNotificationSoundFile)
	if err != nil {
		return
	}
}

func getAudioFile(message string) {
	audioFile, err := getAudio(message)
	if err != nil {
		log.Println(err)
	}
	playVoiceNote(audioFile)
}
