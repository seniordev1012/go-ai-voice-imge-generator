package aigenRest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type Response struct {
	Text string `json:"text"`
	// add other fields as needed
}

func Whisper(pathToFind string) string {
	url := "https://api.openai.com/v1/audio/transcriptions"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open(pathToFind)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("Error closing file: %v", err)
		}
	}(file)
	part1,
		errFile1 := writer.CreateFormFile("file", filepath.Base(pathToFind))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		log.Println(errFile1)
		return ""
	}
	_ = writer.WriteField("model", "whisper-1")
	err := writer.Close()
	if err != nil {
		log.Println(err)
		return ""
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Println(err)
		return ""
	}
	openAi := os.Getenv("OPENAI")
	req.Header.Add("Authorization", "Bearer "+openAi)

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error closing body: %v", err)
		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return ""
	}
	text, err := getTextFromJSON(body)
	if err != nil {
		log.Println(err)
		return ""
	}
	fmt.Println(text)
	return text
}

func getTextFromJSON(rawJSON []byte) (string, error) {
	// unmarshal the JSON data into the struct
	var response Response
	err := json.Unmarshal(rawJSON, &response)
	if err != nil {
		log.Println(err)
		return "", err
	}

	// return the value of the "text" field
	return response.Text, nil
}
