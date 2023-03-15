package main

import (
	"github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/handlers"
	"github.com/hegedustibor/htgo-tts/voices"
	"log"
)

func voiceNote(messageCall string, err error) {
	speakOut := func(speakOut string) string {
		return messageCall
	}

	speech := htgotts.Speech{Folder: "audio", Language: voices.EnglishUK, Handler: &handlers.Native{}}
	err = speech.Speak(speakOut(messageCall))
	log.Println("Ron Swanson says: " + messageCall)
	if err != nil {
		log.Printf("Error speaking: %v", err)

	}
}
