package aigenRest

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

func SpeakOut(innerVoice string) (string, error) {
	var format = ".mp3"
	var audioPath = "voicenotes/"
	innerVoiceLang := "en-US"
	innerVoiceName := "en-US-DavisNeural"

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
		_, out := updateBotChatAudioPath(joinedFileName)
		if out != nil {
			log.Printf("Error updating bot chat audio path: %v", err)
		}

		return joinedFileName, nil
	}
}

func joinFileName(audioPath string, randomString string, format string) string {
	return audioPath + randomString + format
}

func updateBotChatAudioPath(audioPath string) (string, error) {
	// SQL update audio for the last row in the messages table
	dataSourceName := "DB/messages.db"
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Printf("Error opening database: %v", err)
		return "", err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}(db)

	_, err = db.Exec("UPDATE messages SET audio = ? WHERE id = (SELECT id FROM messages ORDER BY id DESC LIMIT 1)", audioPath)
	if err != nil {
		return "", err
	}
	return audioPath, nil
}
