package aigenRest

import (
	"math/rand"
	"strings"
	"time"
)

// Random Name Generator For AI Generated Image
func randomImageName() string {
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
	randomFileName := randomName + ".png"
	imageFileName := "dalleAssets/image_" + randomFileName
	if !strings.HasSuffix(imageFileName, ".png") {
		imageFileName += ".png"
	}
	return imageFileName
}
