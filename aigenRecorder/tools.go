package aigenRecorder

import (
	"math/rand"
	"strings"
	"time"
)

func randomName() string {
	timestamp := time.Now().Format("02-01-2006-15-04-05")
	timestamp = strings.Replace(timestamp, "-", "", -1)
	timestamp = strings.Replace(timestamp, ":", "", -1)

	letterPool := "abcdefghijklmnopqrstuvwxyzABZDEFGHIJKLMNOPQRSTUVWXYZ"
	randomLetter := string(letterPool[rand.Intn(len(letterPool))])

	var letters strings.Builder
	for i := 0; i < 5; i++ {
		letters.WriteByte(letterPool[rand.Intn(len(letterPool))])
	}

	randomName := timestamp + randomLetter + letters.String()
	randomFileName := randomName + ".wav"
	audioFileName := "cache/voice_" + randomFileName
	if !strings.HasSuffix(audioFileName, ".wav") {
		audioFileName += ".wav"
	}
	return audioFileName
}
