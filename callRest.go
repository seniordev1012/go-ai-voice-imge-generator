package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

var outerVoice = "Let me tell you about the Kingdom of Mapungubwe. It was a kingdom in the 13th century." +
	" It was located in the Limpopo River Valley in South Africa and was the first kingdom in Southern Africa." +
	" It was a kingdom of gold and ivory. It was a kingdom of gold and ivory. It was a kingdom of gold and ivory."

func speakOut(innerVoice string) (string, error) {
	var format = ".mp3"
	var audioPath = "voicenotes/"
	innerVoiceLang := "en-US"
	innerVoiceName := "en-US-AriaNeural"

	url := fmt.Sprintf("https://%s.tts.speech.microsoft.com/cognitiveservices/v1", os.Getenv("SPEECH_REGION"))
	method := "POST"
	payload := strings.NewReader(fmt.Sprintf("<speak version='1.0' xmlns='http://www.w3.org/2001/10/synthesis'"+
		" xml:lang='%s'><voice name='%s'>%s</voice></speak>", innerVoiceLang, innerVoiceName, innerVoice))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return "", nil
	}
	os.Getenv("SPEECH_KEY")
	req.Header.Add("Ocp-Apim-Subscription-Key", os.Getenv("SPEECH_KEY"))
	req.Header.Add("Content-Type", "application/ssml+xml")
	req.Header.Add("X-Microsoft-OutputFormat", "audio-16khz-128kbitrate-mono-mp3")
	req.Header.Add("User-Agent", "curl")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}

	generateLetters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	randomString := ""
	for i := 0; i < 10; i++ {
		randomString += string(generateLetters[rand.Intn(len(generateLetters))])
	}
	timestamp := time.Now().Format("2006-01-02-15-04-05")
	timestamp = strings.ReplaceAll(timestamp, "-", "")
	randomString += timestamp
	err = ioutil.WriteFile(audioPath+randomString+format, body, 0644)
	if err != nil {
		fmt.Println(err)
		return "", nil
	} else {

		joinedFileName := joinFileName(audioPath, randomString, format)
		log.Printf("File saved to %s", joinedFileName)
		return joinedFileName, nil
	}
}

func joinFileName(audioPath string, randomString string, format string) string {
	return audioPath + randomString + format
}
